package repository

import (
	"github.com/Sirok47/TOP_GAMES-interfaces-/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Update updates data in mongoDB using "TopGames" structure's object
func (r *TopGamesMongo) Update(g *model.SingleGame) error {
	_, err := r.db.ReplaceOne(
		r.ctx,
		bson.D{
			primitive.E{
				Key:   "_id",
				Value: g.ID,
			},
		},
		bson.M{
			"_id":      g.ID,
			"Name":     g.Name,
			"Rating":   g.Rating,
			"Platform": g.Platform,
			"Date":     g.Date,
		})

	return errors.Wrap(err, "Update failed")
}

// Update updates data in redisDB using "TopGames" structure's object
func (r *TopGamesRedis) Update(g *model.SingleGame) error {
	return r.Create(g)
}

func (r *TopGamesPostgres) Update(g *model.SingleGame) error {
	_, err := r.db.Exec("UPDATE TopGames SET GameName = $1, Rating = $2, Platform = $3, ReleaseDate = $4 WHERE Id = $5", g.Name, g.Rating, g.Platform, g.Date, g.ID)
	return err
}
