package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jessicapaz/kuehne-nagel-challenge/app/models"
)

// HealthCheck returns a successful message if the API is up and running
func HealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			var resp models.HealthCheckResponse
			resp.ApplicationName = "Kuehne + Nagel Challenge"
			byteResp, _ := json.Marshal(resp)

			w.Header().Set("content-type", "application/json")
			w.Write(byteResp)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
