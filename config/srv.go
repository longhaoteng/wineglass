package config

import (
	"strings"

	"github.com/longhaoteng/wineglass/env"
)

var (
	Service = &service{}
)

type service struct {
	Name         string   `json:"name"`
	Version      string   `json:"version"`
	EnablePprof  bool     `json:"enable_pprof"`
	DisableDB    bool     `json:"disable_db"`
	DisableAuth  bool     `json:"disable_auth"`
	DisableRedis bool     `json:"disable_redis"`
	HttpAddr     string   `json:"http_addr"`
	GrpcAddr     string   `json:"grpc_addr"`
	AllowOrigins []string `json:"allow_origins"`
}

func (s *service) init() error {
	enablePprof, err := env.GetBool("ENABLE_PPROF", false)
	if err != nil {
		return err
	}
	disableDB, err := env.GetBool("DISABLE_DB", false)
	if err != nil {
		return err
	}
	disableAuth, err := env.GetBool("DISABLE_AUTH", false)
	if err != nil {
		return err
	}
	disableRedis, err := env.GetBool("DISABLE_REDIS", false)
	if err != nil {
		return err
	}
	Service = &service{
		Version:      env.GetString("VERSION", "latest"),
		EnablePprof:  enablePprof,
		DisableDB:    disableDB,
		DisableAuth:  disableAuth,
		DisableRedis: disableRedis,
		HttpAddr:     env.GetString("HTTP_ADDR", ":8080"),
		GrpcAddr:     env.GetString("GRPC_ADDR", ":50051"),
		AllowOrigins: strings.Split(env.GetString("ALLOW_ORIGINS", "*"), ","),
	}

	return nil
}

func init() {
	AddConfigs(Service)
}
