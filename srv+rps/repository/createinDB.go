package repository

import (
	"github.com/Sirok47/TOP_GAMES-interfaces-/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

// Create inserts data from "TopGames" structure's object into mongoDB
func (r *TopGamesMongo) Create(g *model.SingleGame) error {
	_, err := r.db.InsertOne(
		r.ctx,
		bson.M{
			"_id":      g.ID,
			"Name":     g.Name,
			"Rating":   g.Rating,
			"Platform": g.Platform,
			"Date":     g.Date,
		})

	return errors.Wrap(err, "Create failed")
}

// Create inserts data from "TopGames" structure's object into redisDB
func (r *TopGamesRedis) Create(g *model.SingleGame) error {
	_, err := r.db.Do("HMSET", g.ID, "Name", g.Name, "Rating", g.Rating, "Platform", g.Platform, "Date", g.Date)
	if err != nil {
		return errors.Wrap(err, "Create failed")
	}
	return nil
}

func (r *TopGamesPostgres) Create(g *model.SingleGame) error {
	_, err := r.db.Exec("insert into TopGames (id,GameName,Rating,Platform,ReleaseDate) values ($1,$2,$3,$4,$5)", g.ID, g.Name, g.Rating, g.Platform, g.Date)
	return err
}
