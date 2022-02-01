package grpc

import (
	"context"
	"tinyUrlGRPC/internal/repository"
	"tinyUrlGRPC/pkg/proto"
)

type ServerGrpc struct {
	proto.UnimplementedShorterServer
	r *repository.Reposirory
}

func NewServerGrpc(r *repository.Reposirory) *ServerGrpc {
	return &ServerGrpc{r: r}
}

func (s *ServerGrpc) CreateTinyURL(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {

}

func (s *ServerGrpc) GetTinyURL(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {

}