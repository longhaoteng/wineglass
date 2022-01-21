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
	DB = &dbConf{
		User:         env.GetString("DB_USER", "root"),
		Password:     env.GetString("DB_PASSWORD", ""),
		Host:         env.GetString("DB_HOST", "localhost"),
		Port:         port,
		DBName:       env.GetString("DB_NAME", Service.Name),
		LowThreshold: lowThreshold,
	}

	return nil
}

func init() {
	AddConfigs(DB)
}
