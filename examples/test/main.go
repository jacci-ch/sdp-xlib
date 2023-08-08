package main

import (
	"fmt"
	"github.com/jacci-ch/sdp-xlib/etcdx"
)

func main() {
	rsp, err := etcdx.GetByPrefixAsMap("/sdp/service/endpoints")
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)
}
