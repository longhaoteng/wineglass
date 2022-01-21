package config

import (
	"github.com/longhaoteng/wineglass/consts"
	"github.com/longhaoteng/wineglass/env"
)

var (
	Log = &logConf{}
)

type logConf struct {
	Level string `json:"level"`
}

func (l *logConf) init() error {
	Log = &logConf{
		Level: env.GetString(consts.LogLevel, "info"),
	}

	return nil
}

func init() {
	AddConfigs(Log)
}
