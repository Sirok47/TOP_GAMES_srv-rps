package repository

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete deletes data from mongoDB by ID
func (r *TopGamesMongo) Delete(id int) error {
	_, err := r.db.DeleteOne(
		r.ctx,
		bson.D{
			primitive.E{
				Key:   "_id",
				Value: id,
			},
		})

	return errors.Wrap(err, "Delete failed")
}

// Delete deletes data from redisDB by ID
func (r *TopGamesRedis) Delete(id int) error {
	_, err := r.db.Do("EXPIRE", id, 0)
	return errors.Wrap(err, "Delete failed")
}
