package main

import (
	"fmt"
	"github.com/jacci-ch/sdp-xlib/examples/grpcx/pb"
	"github.com/jacci-ch/sdp-xlib/grpcx"
	"google.golang.org/grpc"
	"sync"
)

func startServer(addr string) {
	listener, err := grpcx.GetListener()
	if err != nil {
		fmt.Printf("err: %v", err)
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterTestServer(server, &TestServer{Name: addr})

	fmt.Println("listen on:", addr)
	if err := server.Serve(listener); err != nil {
		fmt.Printf("err: %v", err)
		panic(err)
	}
}

func main() {
	var group sync.WaitGroup

	//addrs := []string{":9527", ":9528", ":9529"}
	addrs := []string{":9527"}
	for _, addr := range addrs {
		group.Add(1)
		go func(addr string) {
			startServer(addr)
			group.Add(-1)
		}(addr)
	}

	group.Wait()
}
