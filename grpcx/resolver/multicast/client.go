package multicast

import (
	"fmt"
	"google.golang.org/grpc"
	"strings"
)

func Dial(endpoints []string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	target := fmt.Sprintf("%v://%v", Scheme, strings.Join(endpoints, Separator))
	return grpc.Dial(target, opts...)
}
