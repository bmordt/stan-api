package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/bmordt/stan-api/src/services"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	portNum string

	logger *logrus.Entry
)

func init() {
	logger = newLogger()
}

func newLogger() *logrus.Entry {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: false,
	})
	log := logrus.NewEntry(l).WithFields(logrus.Fields{
		"Product": "stan-api",
	})
	return log
}

func main() {
	//put here instead of in init function so test can run
	initEnvVariables()

	muxrouter := mux.NewRouter()

	articleService := services.NewStanService(logger)

	// -- article routes
	muxrouter.HandleFunc("/", articleService.FilterStanJson).Methods("POST")

	//Router end
	logger.Infof("Server listening on port %s", portNum)
	logger.Fatalf("%v", http.ListenAndServe(":"+portNum, muxrouter))
}

//initEnvVariables gets required env variables
func initEnvVariables() {
	portNum = GetAPIPort()
	if strings.Compare(portNum, "") == 0 {
		logger.Fatalf("Server Port env \"APIPORT\" variable is not set: %s", portNum)
	}
}

//GetAPIPort gets the api port from env
func GetAPIPort() string {
	return os.Getenv("APIPORT")
}
