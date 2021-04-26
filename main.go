package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jessicapaz/kuehne-nagel-challenge/app/routes"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HealthCheck()).Methods("GET")
	r.HandleFunc("/files", routes.UploadFile()).Methods("POST")
	r.HandleFunc("/files", routes.ListFiles()).Methods("GET")
	r.HandleFunc("/files/{id}/download", routes.DownloadFile()).Methods("GET")
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Error(err)
	}
}
