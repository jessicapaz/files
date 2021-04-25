package main

import (
	"net/http"
	"os"

	"github.com/jessicapaz/kuehne-nagel-challenge/app/routes"
	log "github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/", routes.HealthCheck())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Error(err)
	}
}
