package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/iannrafisyah/delos/config"
	"github.com/iannrafisyah/delos/migrations"
	"github.com/iannrafisyah/delos/products"
)

var (
	version = flag.Int("version", 0, "Version")
	port    = flag.Int("port", 9080, "Port")
)

func main() {
	flag.Parse()

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Init Config Environment
	config.Environment()

	//Run migration
	if *version > 0 {
		if err := migrations.Migration(&config.Config.PostgreSQL, *version); err != nil {
			fmt.Println(err.Error())
		}
		os.Exit(0)
	}

	//Connection postgreSQL client
	db, err := config.PostgreConnection(&config.Config.PostgreSQL)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	//Run services
	routes := mux.NewRouter()

	products.Routes(routes, config.PostgreConn)

	portStr := strconv.Itoa(*port)
	srv := &http.Server{
		Handler:      routes,
		Addr:         ":" + portStr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	fmt.Println("Server Running in Port : ", portStr)

	log.Fatal(srv.ListenAndServe())
}
