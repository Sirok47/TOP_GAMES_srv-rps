// Package repository works with DB
package repository

import (
	"context"
	"database/sql"

	"github.com/Sirok47/TOP_GAMES-interfaces-/model"
	"github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

// DBTemplate for mongo and redis
type DBTemplate interface {
	Create(g *model.SingleGame) error
	Read(id int) (*model.SingleGame, error)
	Update(g *model.SingleGame) error
	Delete(id int) error
}

// TopGamesMongo stores DB connection's and context's objects for mongoDB
type TopGamesMongo struct {
	db  *mongo.Collection
	ctx context.Context
}

// TopGamesRedis stores DB connection's object for redis
type TopGamesRedis struct {
	db redis.Conn
}

type TopGamesPostgres struct {
	db *sql.DB
}

// NewRepository is a constructor for creating "TopGames"'s object in repository package
func NewMongoRepository(ctx context.Context, dbMongo *mongo.Collection) DBTemplate {
	return &TopGamesMongo{dbMongo, ctx}
}

func NewRedisRepository(dbRedis redis.Conn) DBTemplate {
	return &TopGamesRedis{dbRedis}
}

func NewPostgresRepository(dbPostgres *sql.DB) DBTemplate {
	return &TopGamesPostgres{dbPostgres}
}

