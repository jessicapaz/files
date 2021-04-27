package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jessicapaz/kuehne-nagel-challenge/app/routes"
	"github.com/jessicapaz/kuehne-nagel-challenge/app/services"
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

	fileService := services.NewFileService()
	ttl, _ := strconv.Atoi(os.Getenv("TTL"))
	go func() {
		for {
			fileService.DeleteOldFiles(ttl)
		}
	}()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
