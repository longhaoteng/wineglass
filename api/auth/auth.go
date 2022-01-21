package auth

import "encoding/gob"

const (
	TokenKey = "daydream_token"
)

type User struct {
	ID    int64 `json:"id"`
	State bool  `json:"state"`
}

func init() {
	gob.Register(&User{})
}
