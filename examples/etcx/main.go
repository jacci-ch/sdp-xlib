package main

import (
	"fmt"
	"github.com/jacci-ch/sdp-xlib/etcdx"
	"github.com/jacci-ch/sdp-xlib/jsonx"
)

func main() {
	keys, err := etcdx.GetKeys("/sdp/service/endpoints")
	if err != nil {
		panic(err)
	}

	fmt.Println("keys =", jsonx.Encode(keys))
}
