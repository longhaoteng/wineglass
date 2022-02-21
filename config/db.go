package config

import (
	"github.com/longhaoteng/wineglass/env"
)

var (
	DB = &dbConf{}
)

type dbConf struct {
	User         string `json:"user"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	DBName       string `json:"name"`
	LowThreshold int    `json:"low_threshold"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxLifeTime  int    `json:"max_life_time"`
	MaxIdleTime  int    `json:"max_idle_time"`
}

func (d *dbConf) init() error {
	port, err := env.GetInt("DB_PORT", 3306)
	if err != nil {
		return err
	}
	lowThreshold, err := env.GetInt("DB_LOW_THRESHOLD", 0)
	if err != nil {
		return err
	}
	maxOpenConns, err := env.GetInt("DB_MAX_OPEN_CONNS", 100)
	if err != nil {
		return err
	}
	maxIdleConns, err := env.GetInt("DB_MAX_IDLE_CONNS", 25)
	if err != nil {
		return err
	}
	maxLifeTime, err := env.GetInt("DB_MAX_LIFE_TIME", 0)
	if err != nil {
		return err
	}
	maxIdleTime, err := env.GetInt("DB_MAX_IDLE_TIME", 0)
	if err != nil {
		return err
	}
	DB = &dbConf{
		User:         env.GetString("DB_USER", "root"),
		Password:     env.GetString("DB_PASSWORD", ""),
		Host:         env.GetString("DB_HOST", "localhost"),
		Port:         port,
		DBName:       env.GetString("DB_NAME", Service.Name),
		LowThreshold: lowThreshold,
		MaxOpenConns: maxOpenConns,
		MaxIdleConns: maxIdleConns,
		MaxLifeTime:  maxLifeTime,
		MaxIdleTime:  maxIdleTime,
	}

	return nil
}

func init() {
	AddConfigs(DB)
}
