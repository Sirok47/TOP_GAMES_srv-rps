// Package service just passes values to repository
package service

import (
	"github.com/Sirok47/TOP_GAMES/model"
	"github.com/Sirok47/TOP_GAMES_srv-rps/srv+rps/repository"
)

// TopGames stores DB connection's, context's and next structure's objects for service package
type TopGames struct {
	rps repository.DBTemplate
}

// NewService is a constructor for creating "TopGames"'s object in service package
func NewService(rps repository.DBTemplate) *TopGames {
	return &TopGames{rps}
}

// Read passes id to rps.Read
func (s TopGames) Read(id int) (*model.SingleGame, error) {
	return s.rps.Read(id)
}

// Create passes "TopGames"'s object to rps.Create
func (s TopGames) Create(g *model.SingleGame) error {
	return s.rps.Create(g)
}

// Update passes "TopGames"'s object to rps.Update
func (s TopGames) Update(g *model.SingleGame) error {
	return s.rps.Update(g)
}

// Delete passes id to rps.Delete
func (s TopGames) Delete(id int) error {
	return s.rps.Delete(id)
}
