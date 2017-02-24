package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/csduarte/mattermost-jira/jira"
)

func main() {

	port := os.Getenv("MJ_PORT")
	if port == "" {
		port = "5000"
	}

	addr := os.Getenv("MJ_BIND_ADDRESS")
	if addr == "" {
		addr = "0.0.0.0"
	}

	location := addr + ":" + port
	fmt.Printf("Server starting on %s\n", location)

	jbridge := jira.NewBridge()
	http.HandleFunc("/", jbridge.Handler)

	log.Fatal(http.ListenAndServe(location, nil))
}
