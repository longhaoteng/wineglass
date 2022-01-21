package config

import (
	"strings"

	"github.com/longhaoteng/wineglass/env"
)

var (
	Redis *redisConf
)

type redisConf struct {
	DB        int      `json:"db"`
	Addrs     []string `json:"addrs"`
	Password  string   `json:"password"`
	KeyPrefix string   `json:"prefix"`
}

func (r *redisConf) init() error {
	db, err := env.GetInt("REDIS_DB", 0)
	if err != nil {
		return err
	}
	Redis = &redisConf{
		DB:        db,
		Addrs:     strings.Split(env.GetString("REDIS_ADDRS", "localhost:6379"), ","),
		Password:  env.GetString("REDIS_PASSWORD", ""),
		KeyPrefix: env.GetString("REDIS_PREFIX", Service.Name),
	}

	return nil
}

func init() {
	AddConfigs(Redis)
}
