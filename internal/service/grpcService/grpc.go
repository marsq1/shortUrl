package grpcService

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"
	"tinyUrlGRPC/internal/repository"
	"tinyUrlGRPC/pkg/proto"
)

const (
	salt = "asu13hjsdi1h2jSDAjh23ekhj"
)

type ServerGrpc struct {
	proto.UnimplementedShorterServer
	r *repository.Repository
}

func NewServerGrpc(r *repository.Repository) *ServerGrpc {
	return &ServerGrpc{r: r}
}

func (s *ServerGrpc) CreateTinyURL(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	h := sha1.New()
	h.Write([]byte(req.OriginalUrl + salt))
	bs := h.Sum(nil)
	return &proto.CreateResponse{ShortUrl: fmt.Sprintf("%x", bs[:8])}, nil
}

func (s *ServerGrpc) GetTinyURL(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	return &proto.GetResponse{OriginalUrl: s.r.GetOriginalUrl(req.ShortUrl)}, nil
}

func (s *ServerGrpc) SaveUrl(ctx context.Context, req *proto.SaveRequest) (*proto.SaveResponse, error){
	log.Println(req.ShortUrl, req.OriginalUrl)
	s.r.SaveShortUrl(req.ShortUrl, req.OriginalUrl)
	return &proto.SaveResponse{ShortUrl: req.ShortUrl, OriginalUrl: req.OriginalUrl}, nil
}