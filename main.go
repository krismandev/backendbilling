package main

import (
	"billingdashboard/config"
	"billingdashboard/connections"
	lib "billingdashboard/lib"
	"billingdashboard/routes"
	"bytes"
	"context"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var version = "0.0.0"
var builddate = ""

func main() {
	log.Infof("Running on version %s, build date : %s", version, builddate)

	wg := sync.WaitGroup{}
	closeFlag := make(chan struct{})
	var logFile *os.File
	LoadConfig("config/config.yml", config.Param.Log.FileName, config.Param.Log.Level, logFile)
	conn := connections.InitiateConnections(config.Param)
	defer lib.CloseLog(logFile)
	flag.Parse()
	var err error

	// register route
	routes.InitRoutes(conn, version, builddate)

	lib.RegisterExitSignalHandler(func() {
		log.Info("Shutting down gracefully with exit signal")
		close(closeFlag)
	})
	gracefullShutdownChan := make(chan bool)

	log.Info("System Ready")

	server := &http.Server{Addr: config.Param.ListenPort, Handler: logRequest(http.DefaultServeMux)}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, os.Kill)
	go func() {
		<-exit
		//initiate gracefullShutDown
		log.Info("Initiate gracefully shutdown with exit signal")
		close(gracefullShutdownChan)
		WaitTimeout(&wg, 10*time.Second)
		log.Info("Shutting down gracefully with exit signal")
		server.Shutdown(context.TODO())
	}()
	log.Info(server.ListenAndServe())
	//err = http.ListenAndServe(config.Param.ListenPort, nil)

	if err != nil {
		log.Errorf("Unable to start the server %v", err)
		os.Exit(1)
	}
}

//WaitTimeout to wait with timeout
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()
	select {
	case <-done:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}

//LoadConfig to load config from config.yml file
func LoadConfig(configPath, logFileName, logLevel string, logFile *os.File) {
	configFile := flag.String("config", configPath, "main configuration file")
	config.LoadConfig(configFile)
	flag.Parse()
	log.Infof("Reads configuration from %s", *configFile)
	lib.InitLog(config.Param.Log.FileName, config.Param.Log.Level, logFile)
}

// log incoming request
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Errorf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		var logBody string
		if len(body) > config.Param.MaxBodyLogLength {
			// chunk body log request
			logBody = "(Chunked %d first chars) " + string(body[:config.Param.MaxBodyLogLength])
		} else {
			logBody = string(body)
		}

		// Work / inspect body. You may even modify it!
		// And now set a new body, which will simulate the same data we read:
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		log.WithFields(log.Fields{"ip": r.RemoteAddr, "method": r.Method, "url": r.URL, "body": logBody}).Info("request")
		handler.ServeHTTP(w, r)

	})
}
