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
	b:=a.ID
	for b!=1{
		for i:=2;i<=b;i++{
			if b%i==0{
				a.Name=a.Name+"*"+strconv.Itoa(i)
				b/=i
				break
			}
		}
	}
	a.Name=a.Name+")"
	return a
}
/*module TOP_GAMES-interfaces-

go 1.15

require (
github.com/Sirok47/TOP_GAMES-interfaces- v0.0.0-20201215085411-982006479e4b
github.com/Sirok47/TOP_GAMES_srv-rps v0.0.0-20201215145027-e5ce92972094
github.com/gomodule/redigo v1.8.3
github.com/labstack/echo/v4 v4.1.17
github.com/lib/pq v1.9.0
github.com/pkg/errors v0.9.1
go.mongodb.org/mongo-driver v1.4.4
)*/
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
