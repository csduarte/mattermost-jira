package main

import (
	"flag"
	"net/http"
	"os"
)

func main() {

	var logLocation string
	var logVerbose bool

	flag.StringVar(&logLocation, "log", "./mattermost-jira.log", "Log file path")
	flag.BoolVar(&logVerbose, "v", false, "Sets logs to debug level")
	flag.Parse()

	log := initLog(logLocation, logVerbose)

	port := os.Getenv("MJ_PORT")
	if port == "" {
		port = "5000"
	}

	addr := os.Getenv("MJ_BIND_ADDRESS")
	if addr == "" {
		addr = "0.0.0.0"
	}

	location := addr + ":" + port
	log.Infof("Server starting on %s", location)

	jbridge := bridge.New()
	jbridge.Log = log
	http.HandleFunc("/", jbridge.Handler)

	log.Fatal(http.ListenAndServe(location, nil))
}
