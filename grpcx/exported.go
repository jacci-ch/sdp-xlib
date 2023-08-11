// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package grpcx

import (
	"sync"
)

var (
	gLock   sync.Mutex
	gServer *Server
)

// GetServer
//
// Returns the global server. If the global server is nil this
// function will use the given user-implementation to create
// a new one.
func GetServer(s GrpcServer) (server *Server, err error) {
	if gServer == nil {
		gLock.Lock()
		if gServer == nil {
			gServer, err = NewServer(s)
		}
		gLock.Unlock()
	}

	return gServer, nil
}

// Serve
//
// Start the server with the given implementation.
func Serve(s GrpcServer) (err error) {
	server, err := GetServer(s)
	if err != nil {
		return err
	}

	return server.Serve()
}

// ServeAndWait
//
// Start the server with the given implementation and wait for
// its stopping.
func ServeAndWait(s GrpcServer) error {
	server, err := GetServer(s)
	if err != nil {
		return err
	}

	return server.ServeAndWait()
}

func Stop() {
	if gServer != nil {
		gServer.Stop()
	}
}
