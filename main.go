package main

import (
	"flag"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/iannrafisyah/delos/config"
	"github.com/iannrafisyah/delos/migrations"
	"github.com/iannrafisyah/delos/products"
	"github.com/iannrafisyah/delos/utilities"
	"github.com/sirupsen/logrus"
)

var (
	version = flag.Int("version", 0, "Version")
	port    = flag.Int("port", 9080, "Port")
)

// @title Delos Test API
// @version 2.0
// @description This is a docs products services.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9080
// @BasePath /api/v1
// @query.collection.format multi

func main() {
	flag.Parse()

	//Initiate logs
	utilities.Logs()

	//Load config env from config.yml
	config.Environment()

	//Initiate migration if version value not zero
	if *version > 0 {
		if err := migrations.Migration(&config.Config.PostgreSQL, *version); err != nil {
			utilities.Logger.Panic(err)
		}
		os.Exit(0)
	}

	//Initiate connection postgreSQL client
	db, err := config.PostgreConnection(&config.Config.PostgreSQL)
	if err != nil {
		utilities.Logger.Panic(err)
	}
	defer db.Close()

	//Initiate route and swagger api
	routes := mux.NewRouter()
	routes.Use(utilities.LogsMiddleware)

	//Run services products
	products.Routes(routes, config.PostgreConn)

	//Start server
	portStr := strconv.Itoa(*port)
	srv := &http.Server{
		Handler:      routes,
		Addr:         ":" + portStr,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	logrus.Infof("Server Running in Port : ", portStr)
	utilities.Logger.Fatal(srv.ListenAndServe())
}
