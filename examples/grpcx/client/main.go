package main

import (
	"context"
	"fmt"
	"github.com/jacci-ch/sdp-xlib/examples/grpcx/pb"
	_ "github.com/jacci-ch/sdp-xlib/grpcx/resolver/multicast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
	"time"
)

func main() {
	conn, err := grpc.Dial(
		"multicast://localhost:9527,localhost:9528,localhost:9529",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"LoadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		panic(err)
	}

	testClient := pb.NewTestClient(conn)
	for cc := 0; cc < 1000; cc++ {
		rsp, err := testClient.Greet(context.Background(), &pb.Request{Name: strconv.FormatInt(int64(cc), 10)})
		if err != nil {
			panic(err)
		}

		fmt.Printf("rsp  = %v\n", rsp)
		time.Sleep(3 * time.Second)
	}
}
