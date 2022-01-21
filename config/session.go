package config

import (
	"github.com/longhaoteng/wineglass/consts"
	"github.com/longhaoteng/wineglass/env"
)

var (
	Session = &sessionConf{}
)

type sessionConf struct {
	Store    string `json:"store"`
	Secret   string `json:"secret"`
	MaxAge   int    `json:"max_age"`
	HttpOnly bool   `json:"http_only"`
	DB       int    `json:"db"`
}

func (s *sessionConf) init() error {
	maxAge, err := env.GetInt("SESSION_MAX_AGE", 604800)
	if err != nil {
		return err
	}
	httpOnly, err := env.GetBool("SESSION_HTTP_ONLY", false)
	if err != nil {
		return err
	}
	db, err := env.GetInt("SESSION_DB", 0)
	if err != nil {
		return err
	}
	Session = &sessionConf{
		Store:    env.GetString("SESSION_STORE", consts.MemoryStore),
		Secret:   env.GetString("SESSION_SECRET", "wineglass"),
		MaxAge:   maxAge,
		HttpOnly: httpOnly,
		DB:       db,
	}

	return nil
}

func init() {
	AddConfigs(Session)
}
