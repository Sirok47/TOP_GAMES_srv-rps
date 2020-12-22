package main

import (
	"database/sql"
	"fmt"
	grpcpb "github.com/Sirok47/TOP_GAMES-interfaces-/grpc"
	"github.com/Sirok47/TOP_GAMES_srv-rps/srv+rps/repository"
	"github.com/Sirok47/TOP_GAMES_srv-rps/srv+rps/service"
	"github.com/gomodule/redigo/redis"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"time"
)
func main() {
	const (
		dbChoice = 0 // 0 for Postgres, 1 for Mongo, 2 for Redis
		timeout  = 10
	)
	var (
		ctx        context.Context   = nil
		collection *mongo.Collection = nil
		conn       redis.Conn        = nil
		err        error
		con *service.TopGames
	)
	db, err := sql.Open("postgres", "user=postgres password=glazirovanniisirok dbname=TOP_GAMES sslmode=disable")
	if err!= nil{fmt.Println(err)}
	con = service.NewService(repository.NewPostgresRepository(db),db)
	switch dbChoice {
		case 0:
		case 1:
			var cancel context.CancelFunc
			client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
			ctx, cancel = context.WithTimeout(context.Background(), timeout*time.Second)

			defer cancel()

			_ = client.Connect(ctx)
			collection = client.Database("TOP_GAMES").Collection("TopGames")
			ctx = context.TODO()
			con = service.NewService(repository.NewMongoRepository(ctx, collection),db)
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
			con = service.NewService(repository.NewRedisRepository(conn),db)
		default:
			return
		}

s:=grpc.NewServer()
grpcpb.RegisterCRUDServer(s,con)
l,_:=net.Listen("tcp",":8080")
s.Serve(l)
}