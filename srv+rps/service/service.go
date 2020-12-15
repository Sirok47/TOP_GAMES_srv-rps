// Package service just passes values to repository
package service

import (
	"github.com/Sirok47/TOP_GAMES-interfaces-/model"
	"github.com/Sirok47/TOP_GAMES_srv-rps/srv+rps/repository"
	"strconv"
)

// TopGames stores DB connection's, context's and next structure's objects for service package
type TopGames struct {
	rps repository.DBTemplate
}

func simpDigits(a *model.SingleGame) *model.SingleGame{
	a.Name=a.Name+"(ID=1"
	for a.ID!=1{
		for i:=2;i<=a.ID;i++{
			if a.ID%i==0{
				a.Name=a.Name+"+"+strconv.Itoa(i)
				a.ID/=i
				break
			}
		}
	}
	a.Name=a.Name+")"
	return a
}

// NewService is a constructor for creating "TopGames"'s object in service package
func NewService(rps repository.DBTemplate) *TopGames {
	return &TopGames{rps}
}

// Read passes id to rps.Read
func (s TopGames) Read(id int) (*model.SingleGame, error) {
	g,err:=s.rps.Read(id)
	g=simpDigits(g)
	return g,err
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
