package repositories

import "context"

func (r *repo) UpdateOne(data, filter interface{}, tableName string) error {
	_, err := r.db.Collection(tableName).UpdateOne(context.TODO(), filter, data)
	if err != nil {
		return err
	}
	return nil
}
