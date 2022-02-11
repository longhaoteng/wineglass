package auth

import (
	_ "embed"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	scas "github.com/qiangmzsx/string-adapter/v2"

	"github.com/longhaoteng/wineglass/config"
	"github.com/longhaoteng/wineglass/db"
)

const (
	ruleDBName    = "auth"
	ruleTableName = "rule"
)

var (
	//go:embed model.conf
	conf string
	//go:embed policy.csv
	policy string
	e      *casbin.Enforcer
)

func Init() error {
	m, err := model.NewModelFromString(conf)
	if err != nil {
		return err
	}
	if config.Service.DisableAuth {
		e, err = casbin.NewEnforcer(m, scas.NewAdapter(policy))
		if err != nil {
			return err
		}
	} else {
		a, err := gormadapter.NewAdapterByDBUseTableName(db.DB, ruleDBName, ruleTableName)
		if err != nil {
			return err
		}
		e, err = casbin.NewEnforcer(m, a)
		if err != nil {
			return err
		}
	}

	if err = e.LoadPolicy(); err != nil {
		return err
	}

	return nil
}

func Enforce(vs ...interface{}) (bool, error) {
	return e.Enforce(vs...)
}

func LoadPolicy() error {
	return e.LoadPolicy()
}

func SavePolicy() error {
	return e.SavePolicy()
}
