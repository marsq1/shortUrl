package main

import (
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
	"tinyUrlGRPC/internal/delivery/handle"
	"tinyUrlGRPC/internal/repository"
	"tinyUrlGRPC/internal/service/grpcService"
	"tinyUrlGRPC/pkg/proto"
)

const (
	portRest = "8000"
	portGrpc = "9000"
)

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
	Host: "localhost",
	Port: "5432",
	Username: "postgres",
	Password: "qwerty",
	DBName: "grpcShort",
	SSLMode: "disable",
	})
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	s := grpc.NewServer()

	repos := repository.NewRepository(db)
	serviceGrpc := grpcService.NewServerGrpc(repos)
	proto.RegisterShorterServer(s, serviceGrpc)
	l, err := net.Listen("tcp", ":"+portGrpc)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	go func() {
		if err := s.Serve(l); err != nil {
			log.Fatal(err.Error())
			return
		}
	}()

	h := handle.NewHandle(serviceGrpc, portRest)
	h.StartServer()
}