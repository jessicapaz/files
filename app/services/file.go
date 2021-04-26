package services

import (
	"context"
	"log"

	"github.com/jessicapaz/kuehne-nagel-challenge/app/db"
	"github.com/jessicapaz/kuehne-nagel-challenge/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FileService contains all file actions
type FileService struct{}

// NewFileService creates a new FileService
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

// List list all files
func (fs *FileService) List() ([]*models.File, error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}

	collection := client.Database("challenge").Collection("files")

	findOptions := options.Find()
	findOptions.SetLimit(2)

	var results []*models.File

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		var file models.File
		err := cur.Decode(&file)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &file)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	return results, nil
}

// Get gets a file by id
func (fs *FileService) Get(id string) (*models.File, error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}

	collection := client.Database("challenge").Collection("files")

	var result *models.File

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: oid}}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}
	return result, err
}
