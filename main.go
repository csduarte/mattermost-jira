package main

import (
	"flag"
	"net/http"

	"github.com/csduarte/mattermost-jira/bridge"
)

func main() {

	var logLocation, addr, port string
	var logVerbose bool

	flag.StringVar(&logLocation, "log", "./mattermost-jira.log", "Log file path")
	flag.BoolVar(&logVerbose, "v", false, "Sets logs to debug level")
	flag.StringVar(&addr, "address", "0.0.0.0", "Bind adress")
	flag.StringVar(&port, "port", "5000", "Listening port")
	flag.Parse()

	log := initLog(logLocation, logVerbose)
	log.Info("Log Location:", logLocation)
	log.Info("Verbose Log:", logVerbose)

	location := addr + ":" + port
	log.Infof("Server starting on %s", location)

	jbridge := bridge.New()
	jbridge.Log = log
	http.HandleFunc("/", jbridge.Handler)

	log.Fatal(http.ListenAndServe(location, nil))
}
