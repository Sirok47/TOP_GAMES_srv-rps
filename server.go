package main

import (
	grpcpb "github.com/Sirok47/TOP_GAMES-interfaces-/grpc"
	"github.com/Sirok47/TOP_GAMES_srv-rps/srv+rps/service"
	"google.golang.org/grpc"
	"net"
)

func main() {
s:=grpc.NewServer()
srv := &service.TopGames{}
grpcpb.RegisterCRUDServer(s,srv)
l,_:=net.Listen("tcp",":8080")
s.Serve(l)
}