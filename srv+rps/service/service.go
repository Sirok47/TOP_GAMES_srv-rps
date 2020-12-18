// Package service just passes values to repository
package service

import (
	"context"
	"database/sql"
	"fmt"
	grpcpb "github.com/Sirok47/TOP_GAMES-interfaces-/grpc"
	"github.com/Sirok47/TOP_GAMES-interfaces-/model"
	"github.com/Sirok47/TOP_GAMES_srv-rps/srv+rps/repository"
	"github.com/gomodule/redigo/redis"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
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
const (
	dbChoice = 0 // 0 for Postgres, 1 for Mongo, 2 for Redis
	timeout  = 10
)
var (
	ctx        context.Context   = nil
	collection *mongo.Collection = nil
	conn       redis.Conn        = nil
	err        error
	con *TopGames
)
func init() {
	switch dbChoice {
	case 0:
		db, err := sql.Open("postgres", "user=postgres password=glazirovanniisirok dbname=TOP_GAMES sslmode=disable")
		if err!= nil{fmt.Println(err)}
		con = NewService(repository.NewPostgresRepository(db))
	case 1:
		var cancel context.CancelFunc
		client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
		ctx, cancel = context.WithTimeout(context.Background(), timeout*time.Second)

		defer cancel()

		_ = client.Connect(ctx)
		collection = client.Database("TOP_GAMES").Collection("TopGames")
		ctx = context.TODO()
		con = NewService(repository.NewMongoRepository(ctx, collection))
	case 2:
		conn, err = redis.Dial("tcp", "localhost:6379")
		if err != nil {
			return
		}
		defer func() {
			err := conn.Close()
			if err != nil {
				return
			}
		}()
		con = NewService(repository.NewRedisRepository(conn))
	default:
		return
	}
}

// Read passes id to rps.Read
func (s *TopGames) Read(ctx context.Context, rqs *grpcpb.Id) (*grpcpb.Structmsg, error) {
	g,err:=(TopGames{con.rps}).rps.Read(int(rqs.ID))
	simpDigits(g)
	err2:=""
	if err != nil{
		err2=err.Error()
	}
	gg:=&grpcpb.Structmsg{ID: int32(g.ID),Name: g.Name,Rating: int32(g.Rating),Platform: g.Platform,Date: g.Date,Err: err2}
	return gg,nil
}
// Create passes "TopGames"'s object to rps.Create
func (*TopGames) Create(ctx context.Context,g *grpcpb.Structmsg) (*grpcpb.Errmsg,error) {
	gg:=&model.SingleGame{ID: int(g.ID), Name: g.Name, Rating: float64(g.Rating), Platform: g.Platform, Date: g.Date}
	err:=(TopGames{con.rps}).rps.Create(gg)
	err2:=""
	if err != nil{
		err2=err.Error()
	}
	return &grpcpb.Errmsg{Err: err2},nil
}

// Update passes "TopGames"'s object to rps.Update
func (s *TopGames) Update(ctx context.Context,g *grpcpb.Structmsg) (*grpcpb.Errmsg,error) {
	gg:=&model.SingleGame{ID: int(g.ID), Name: g.Name, Rating: float64(g.Rating), Platform: g.Platform, Date: g.Date}
	err:=(TopGames{con.rps}).rps.Create(gg)
	err2:=""
	if err != nil{
		err2=err.Error()
	}
	return &grpcpb.Errmsg{Err: err2},nil
}

// Delete passes id to rps.Delete
func (s *TopGames) Delete(ctx context.Context,rqs *grpcpb.Id) (*grpcpb.Errmsg,error) {
	err:=(TopGames{con.rps}).rps.Delete(int(rqs.ID))
	err2:=""
	if err != nil{
		err2=err.Error()
	}
	return &grpcpb.Errmsg{Err: err2},nil
}

