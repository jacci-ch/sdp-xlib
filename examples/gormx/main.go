package main

import (
	"github.com/jacci-ch/sdp-xlib/gormx"
	"github.com/jacci-ch/sdp-xlib/jsonx"
	"github.com/jacci-ch/sdp-xlib/logx"
)

type User struct {
	Id       int64
	Username string
}

func main() {
	user := &User{}

	if tx := gormx.DB.Table("sdp_user").First(&user, 1); tx.Error != nil {
		panic(tx.Error)
	}

	logx.Logger.Info("user =", jsonx.Encode(user))
}
