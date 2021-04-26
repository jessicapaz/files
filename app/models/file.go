package models

import "time"

// File contains the file info
type File struct {
	Blob      []byte    `json:"blob"`
	CreatedAt time.Time `json:"createdAt"`
}
