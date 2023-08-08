package main

import (
	"context"
	"fmt"
	"github.com/jacci-ch/sdp-xlib/examples/grpcx/pb"
)

type TestServer struct {
	pb.UnimplementedTestServer

	Name string
}

func (s *TestServer) Greet(ctx context.Context, args *pb.Request) (*pb.Response, error) {
	fmt.Printf("%v: %v\n", s.Name, args.Name)
	return &pb.Response{Reply: fmt.Sprintf("%v reply ok", s.Name)}, nil
}
