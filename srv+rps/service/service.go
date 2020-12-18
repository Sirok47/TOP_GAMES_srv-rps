// Package service just passes values to repository
package service

import (
	"context"
	grpcpb "github.com/Sirok47/TOP_GAMES-interfaces-/grpc"
	"github.com/Sirok47/TOP_GAMES-interfaces-/model"
	"github.com/Sirok47/TOP_GAMES_srv-rps/srv+rps/repository"
	_ "github.com/lib/pq"
	"strconv"
)

// TopGames stores DB connection's, context's and next structure's objects for service package
type TopGames struct {
	rps repository.DBTemplate
}

func simpDigits(a *model.SingleGame) {
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
}
// NewService is a constructor for creating "TopGames"'s object in service package
func NewService(rps repository.DBTemplate) *TopGames {
	return &TopGames{rps}
}


// Read passes id to rps.Read
func (s *TopGames) Read(ctx context.Context, rqs *grpcpb.Id) (*grpcpb.Structmsg, error) {
	g,err:=s.rps.Read(int(rqs.ID))
	simpDigits(g)
	err2:=""
	if err != nil{
		err2=err.Error()
	}
	gg:=&grpcpb.Structmsg{ID: int32(g.ID),Name: g.Name,Rating: int32(g.Rating),Platform: g.Platform,Date: g.Date,Err: err2}
	return gg,nil
}
// Create passes "TopGames"'s object to rps.Create
func (s *TopGames) Create(ctx context.Context,g *grpcpb.Structmsg) (*grpcpb.Errmsg,error) {
	gg:=&model.SingleGame{ID: int(g.ID), Name: g.Name, Rating: float64(g.Rating), Platform: g.Platform, Date: g.Date}
	err:=s.rps.Create(gg)
	err2:=""
	if err != nil{
		err2=err.Error()
	}
	return &grpcpb.Errmsg{Err: err2},nil
}

// Update passes "TopGames"'s object to rps.Update
func (s *TopGames) Update(ctx context.Context,g *grpcpb.Structmsg) (*grpcpb.Errmsg,error) {
	gg:=&model.SingleGame{ID: int(g.ID), Name: g.Name, Rating: float64(g.Rating), Platform: g.Platform, Date: g.Date}
	err:=s.rps.Create(gg)
	err2:=""
	if err != nil{
		err2=err.Error()
	}
	return &grpcpb.Errmsg{Err: err2},nil
}

// Delete passes id to rps.Delete
func (s *TopGames) Delete(ctx context.Context,rqs *grpcpb.Id) (*grpcpb.Errmsg,error) {
	err:=s.rps.Delete(int(rqs.ID))
	err2:=""
	if err != nil{
		err2=err.Error()
	}
	return &grpcpb.Errmsg{Err: err2},nil
}

