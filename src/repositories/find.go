package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) FindAll(where interface{}, tableName string, opt *options.FindOptions) (*mongo.Cursor, error) {
	result, err := r.db.Collection(tableName).Find(context.TODO(), where, opt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repo) FindOne(res interface{}, where interface{}, tableName string) error {
	result := r.db.Collection(tableName).FindOne(context.TODO(), where)
	if result.Err() != nil {
		return result.Err()
	}

	result.Decode(res)
	return nil
}
