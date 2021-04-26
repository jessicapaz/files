package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jessicapaz/kuehne-nagel-challenge/app/models"
	"github.com/jessicapaz/kuehne-nagel-challenge/app/services"
)

// UploadFile returns a successful message if the file have been successfully uploaded
func UploadFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			r.ParseMultipartForm(10 << 20)

			file, _, err := r.FormFile("file")

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
			fileModel.Blob = fileBytes
			fileModel.CreatedAt = time.Now()
			err = fileService.Upload(&fileModel)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusCreated)
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
