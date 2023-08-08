package main

import (
	"fmt"
	_ "github.com/jacci-ch/sdp-xlib/cfgx"
	"github.com/jacci-ch/sdp-xlib/logx"
	"strings"
	"time"
)

func main() {
	sb := strings.Builder{}

	for cc := 0; cc < 1000; cc++ {
		sb.WriteString("test string ")
	}

	for cc := 0; cc < 100000; cc++ {
		fmt.Println("cc =", cc)
		logx.Logger.Info(sb.String())
		time.Sleep(100 * time.Millisecond)
	}
}
