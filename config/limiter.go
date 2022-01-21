package config

import (
	"github.com/longhaoteng/wineglass/consts"
	"github.com/longhaoteng/wineglass/env"
)

var (
	Limiter = &limiterConf{}
)

type limiterConf struct {
	Store string `json:"store"`
	// format:<limit>-<period>
	// 5 reqs/second: "5-S"
	// 10 reqs/minute: "10-M"
	// 1000 reqs/hour: "1000-H"
	// 2000 reqs/day: "2000-D"
	Limit string `json:"limit"`
}

func (l *limiterConf) init() error {
	Limiter = &limiterConf{
		Store: env.GetString("LIMITER_STORE", consts.MemoryStore),
		Limit: env.GetString("LIMITER_LIMIT", "10-S"),
	}

	return nil
}

func init() {
	AddConfigs(Limiter)
}
