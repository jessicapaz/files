package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jessicapaz/files/app/models"
	"github.com/jessicapaz/files/app/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UploadFile returns a successful message if the file have been successfully uploaded
func UploadFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			r.ParseMultipartForm(10 << 20)

			file, handler, err := r.FormFile("file")

			var e models.ErrorResponse
			if err != nil {
				e.Error = err.Error()
				byteResp, _ := json.Marshal(e)

				w.WriteHeader(http.StatusBadRequest)
				w.Write(byteResp)
			}
			defer file.Close()

			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				e.Error = err.Error()
				byteResp, _ := json.Marshal(e)

				w.WriteHeader(http.StatusBadRequest)
				w.Write(byteResp)
			}

			fileService := services.NewFileService()

			var fileModel models.File
			fileModel.ID = primitive.NewObjectID()
			fileModel.Name = handler.Filename
			fileModel.Blob = fileBytes
			fileModel.CreatedAt = time.Now()

			err = fileService.Upload(&fileModel)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				var fileResp models.FileResponse
				fileResp.ID = fileModel.ID.Hex()
				fileResp.Name = fileModel.Name
				fileResp.CreatedAt = fileModel.CreatedAt
				fileResp.URL = fileResp.BuildURL(r, fileResp.ID)

				byteResp, _ := json.Marshal(fileResp)
				w.WriteHeader(http.StatusCreated)
				w.Write(byteResp)
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

// DownloadFile downloads a file
func DownloadFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		fileService := services.NewFileService()
		file, _ := fileService.Get(vars["id"])

		contentDisposition := fmt.Sprintf("attachment; filename=%s", file.Name)
		w.Header().Set("Content-Disposition", contentDisposition)
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

		reader := bytes.NewReader(file.Blob)

		io.Copy(w, reader)
	}
}

// ListFiles list all files
func ListFiles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileService := services.NewFileService()
		files, _ := fileService.List()

		var resp []models.FileResponse
		for _, file := range files {
			var fileResp models.FileResponse
			fileResp.ID = file.ID.Hex()
			fileResp.Name = file.Name
			fileResp.CreatedAt = file.CreatedAt

			fileResp.URL = fileResp.BuildURL(r, fileResp.ID)
			resp = append(resp, fileResp)
		}

		byteResp, _ := json.Marshal(resp)
		w.Header().Set("content-type", "application/json")
		w.Write(byteResp)
	}
}
