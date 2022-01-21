package server

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/longhaoteng/wineglass/config"
	log "github.com/longhaoteng/wineglass/logger"
)

var (
	grpcSrv = &grpcEntry{}
)

type grpcEntry struct {
	g       *grpc.Server
	handles []Handle
}

type Handle2 struct {
	SD      *grpc.ServiceDesc
	Service interface{}
}

type Handle func(*grpc.Server)

// AddHandles run before service starts
func AddHandles(srvHandles ...Handle) {
	grpcSrv.handles = append(grpcSrv.handles, srvHandles...)
}

func (g *grpcEntry) init(opts ...Option) error {
	grpcSrv.g = grpc.NewServer()

	for _, handle := range g.handles {
		handle(grpcSrv.g)
	}

	// Register reflection service on gRPC server
	reflection.Register(grpcSrv.g)

	return nil
}

func (g *grpcEntry) run() error {
	lis, err := net.Listen("tcp", config.Service.GrpcAddr)
	if err != nil {
		return err
	}

	if log.V(log.InfoLevel) {
		log.Logf(log.InfoLevel, "GRPC Listening on %s", config.Service.GrpcAddr)
	}

	if err = grpcSrv.g.Serve(lis); err != nil {
		panic(err)
	}

	return nil
}

func (g *grpcEntry) stop() error {
	g.g.Stop()
	return nil
}
