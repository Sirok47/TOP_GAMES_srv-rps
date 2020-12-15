package repository

import (
	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Read gets data from mongoDB by ID and writes it into "TopGames"'s object
func (r *TopGamesMongo) Read(id int) (*model.SingleGame, error) {
	g := &model.SingleGame{ID: id, Name: "---", Rating: 0, Platform: "---", Date: "---"}
	err := r.db.FindOne(
		r.ctx,
		bson.D{
			primitive.E{
				Key:   "_id",
				Value: g.ID,
			},
		}).Decode(&g)

	return g, errors.Wrap(err, "Read failed")
}

// Read gets data from redisDB by ID and writes it into "TopGames"'s object
func (r *TopGamesRedis) Read(id int) (*model.SingleGame, error) {
	var err error
	g := &model.SingleGame{ID: id, Name: "---", Rating: 0, Platform: "---", Date: "---"}
	g.Name, err = redis.String(r.db.Do("HGET", g.ID, "Name"))
	if err != nil {
		return nil, errors.Wrap(err, "Read failed")
	}
	g.Rating, err = redis.Float64(r.db.Do("HGET", g.ID, "Rating"))
	if err != nil {
		return nil, errors.Wrap(err, "Read failed")
	}
	g.Platform, err = redis.String(r.db.Do("HGET", g.ID, "Platform"))
	if err != nil {
		return nil, errors.Wrap(err, "Read failed")
	}
	g.Date, err = redis.String(r.db.Do("HGET", g.ID, "Date"))
	if err != nil {
		return nil, errors.Wrap(err, "Read failed")
	}
	return g, nil
}
