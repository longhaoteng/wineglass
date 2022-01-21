package config

import (
	"github.com/longhaoteng/wineglass/consts"
	"github.com/longhaoteng/wineglass/env"
)

var (
	Env   string
	confs []Conf
)

type Conf interface {
	init() error
}

func AddConfigs(cs ...Conf) {
	confs = append(confs, cs...)
}

func init() {
	Env = env.GetString(consts.Env, consts.Prod)
}

func Init() error {
	for _, conf := range confs {
		if err := conf.init(); err != nil {
			return err
		}
	}
	return nil
}

func IsDevEnv() bool {
	return Env == consts.Dev
}

func IsTestEnv() bool {
	return Env == consts.Test
}

func IsProdEnv() bool {
	return Env == consts.Prod || Env == consts.Release
}
