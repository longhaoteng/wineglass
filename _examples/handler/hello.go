package handler

import (
	"context"

	"google.golang.org/grpc"

	"github.com/longhaoteng/wineglass/_examples/proto"
	log "github.com/longhaoteng/wineglass/logger"
	"github.com/longhaoteng/wineglass/server"
)

type HelloServer struct{}

func (h *HelloServer) Hello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	log.Field("name", in.GetName()).Log(log.InfoLevel, "received")
	return &proto.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func init() {
	server.AddHandles(func(g *grpc.Server) {
		proto.RegisterHelloServer(g, &HelloServer{})
	})
}
