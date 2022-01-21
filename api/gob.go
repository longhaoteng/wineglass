package api

import "encoding/gob"

func AddGobModels(models ...interface{}) {
	for _, m := range models {
		gob.Register(m)
	}
}
