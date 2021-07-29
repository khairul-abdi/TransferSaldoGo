package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repo struct {
	db *mongo.Database
}

type Repo interface {
	FindAll(where interface{}, tableName string, opt *options.FindOptions) (*mongo.Cursor, error)
	FindOne(res interface{}, where interface{}, tableName string) error
	InsertOne(data interface{}, tableName string) error
	UpdateOne(data, filter interface{}, tableName string) error
}

func NewRepo(db *mongo.Database) Repo {
	return &repo{db: db}
}
