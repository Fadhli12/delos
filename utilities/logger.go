package utilities

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	Logger = logrus.New()
)

const (
	dir     = "logs"
	errFile = "error.log"
)

func Logs() {
	//If logrus.log not exist, will automatically create file log
	if _, err := os.Stat(fmt.Sprintf("%v/%v", dir, errFile)); os.IsNotExist(err) {
		file, err := os.Create(fmt.Sprintf("%v/%v", dir, errFile))
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}

	file, err := os.OpenFile(fmt.Sprintf("%v/%v", dir, errFile), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	Logger.Level = logrus.ErrorLevel
	Logger.Level = logrus.PanicLevel
	Logger.Formatter = &logrus.TextFormatter{}
	Logger.Out = file
}

func LogsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
		next.ServeHTTP(w, r)
	})
}
