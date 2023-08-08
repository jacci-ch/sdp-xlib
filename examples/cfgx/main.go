package main

import (
	_ "github.com/jacci-ch/sdp-xlib/cfgx"
	"github.com/jacci-ch/sdp-xlib/logx"
)

func main() {
	logx.Logger.Fatal("test fatal error")
	logx.Logger.Info("hello world")
}
