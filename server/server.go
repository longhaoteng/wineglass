package server

import (
	"context"
	"os"
	"os/signal"

	"github.com/longhaoteng/wineglass/api/auth"
	"github.com/longhaoteng/wineglass/cache/redis"
	"github.com/longhaoteng/wineglass/config"
	"github.com/longhaoteng/wineglass/cron"
	"github.com/longhaoteng/wineglass/cron/jobs"
	_ "github.com/longhaoteng/wineglass/cron/jobs"
	"github.com/longhaoteng/wineglass/db"
	"github.com/longhaoteng/wineglass/logger"
	signalutil "github.com/longhaoteng/wineglass/utils/signal"
)

var (
	srv *server
)

type Server interface {
	init(opts ...Option) error
	run() error
	stop() error
}

type server struct {
	api  *apiEntry
	grpc *grpcEntry
	opts *Options
}

func Init(opts ...Option) {
	srv = &server{
		api:  api,
		grpc: grpcSrv,
		opts: &Options{
			Context: context.Background(),
		},
	}

	if err := config.Init(); err != nil {
		panic(err)
	}

	opts = append(opts,
		BeforeStart(func() error {
			cron.AddJobs(&jobs.CasbinPolicy{})
			return auth.Init()
		}),
		AfterStart(func() error { return cron.Init() }),
		BeforeStop(
			func() error {
				cron.Stop()
				return nil
			},
		),
	)

	for _, o := range opts {
		o(srv.opts)
	}

	if err := logger.Init(); err != nil {
		panic(err)
	}

	if !config.Service.DisableDB {
		if err := db.Init(); err != nil {
			panic(err)
		}
	}

	if !config.Service.DisableRedis {
		if err := redis.Init(); err != nil {
			panic(err)
		}
	}

	if err := srv.grpc.init(opts...); err != nil {
		panic(err)
	}

	if err := srv.api.init(opts...); err != nil {
		panic(err)
	}

}

func Run() {
	for _, fn := range srv.opts.BeforeStart {
		if err := fn(); err != nil {
			panic(err)
		}
	}

	go func() {
		if err := srv.grpc.run(); err != nil {
			panic(err)
		}
	}()
	go func() {
		if err := srv.api.run(); err != nil {
			panic(err)
		}
	}()

	for _, fn := range srv.opts.AfterStart {
		if err := fn(); err != nil {
			panic(err)
		}
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signalutil.Shutdown()...)

	select {
	// wait on kill signal
	case <-ch:
	// wait on context cancel
	case <-srv.opts.Context.Done():
	}

	for _, fn := range srv.opts.BeforeStop {
		if err := fn(); err != nil {
			panic(err)
		}
	}

	if err := srv.grpc.stop(); err != nil {
		panic(err)
	}
	if err := srv.api.stop(); err != nil {
		panic(err)
	}

	for _, fn := range srv.opts.AfterStop {
		if err := fn(); err != nil {
			panic(err)
		}
	}
}
