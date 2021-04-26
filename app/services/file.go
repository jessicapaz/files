package services

import (
	"context"

	"github.com/jessicapaz/kuehne-nagel-challenge/app/db"
	"github.com/jessicapaz/kuehne-nagel-challenge/app/models"
)

type FileServiceInterface interface {
	Upload(file *models.File) error
}

type FileService struct{}

// NewFileService creates a new fileService
func NewFileService() *FileService {
	return &FileService{}
}

// Upload uploads a file
func (fs *FileService) Upload(file *models.File) error {
	client, err := db.Client()
	if err != nil {
		return err
	}
	collection := client.Database("challenge").Collection("files")

	_, err = collection.InsertOne(context.TODO(), &file)
	if err != nil {
		return err
	}
	return nil
}
