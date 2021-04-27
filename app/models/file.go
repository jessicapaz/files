package models

import (
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// File contains the file info
type File struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"fileName"`
	Blob      []byte             `json:"blob"`
	CreatedAt time.Time          `json:"createdAt"`
}

// FileResponse contains the file info
type FileResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"fileName"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
}

// BuildURL creates a new file URL
func (fr *FileResponse) BuildURL(request *http.Request, fileID string) string {
	URLScheme := "https"
	if request.TLS == nil {
		URLScheme = "http"
	}
	return fmt.Sprintf("%s://%s/files/%s/download", URLScheme, request.Host, fileID)
}
