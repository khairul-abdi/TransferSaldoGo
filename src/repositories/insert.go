package repositories

import "context"

func (r *repo) InsertOne(data interface{}, tableName string) error {
	_, err := r.db.Collection(tableName).InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}
