// Copyright 2023 to now() The SDP Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpcx

import (
	"errors"
	"sync"

	"github.com/jacci-ch/sdp-xlib/logx"
	"google.golang.org/grpc"
)

type Server struct {
	// Configuration used to create real grpc server, includes listen addr
	// and port, NATed addr and port.
	//
	// In fact only the creation of server listener is depends on these
	// configurations.
	cfg *Config

	// A flag indicates that the server's status. This flag will be set
	// to be true before grpc.Serve() is called, cause this function will
	// block the go routine.
	//
	// This flag will be set to be false after grpc.Serve() exit or execution
	// of function Server.Stop.
	isStart bool

	// A sync instance to wait for serve routine to exit. We can call the method
	// Server.ServeAndWait() to wait for it.
	waitGroup sync.WaitGroup

	// The real grpc.Server object create with grpc.NewServer().
	realServer *grpc.Server

	// A list of ser-registered grpc server. GrpcServer.RegisterRpc method
	// will be called before serve.
	servers []GrpcServer
}

// Serve - start the server. The BeforeStartHook of all registered servers
// will be called one-by-one before server is started. And the AfterStartHook
// hooks will be called after the server is started.
func (s *Server) Serve() error {
	if s.isStart {
		return errors.New("grpcx: server is already started")
	} else {
		// executes all BeforeStartHook hooks.
		if err := s.execBeforeStartHook(); err != nil {
			logx.Errorf("grpcx: %v", err)
			return err
		}

		if err := s.start(); err != nil {
			logx.Errorf("grpcx: %v", err)
			return err
		}

		// executes all AfterStartHook hooks.
		if err := s.execAfterStartHook(); err != nil {
			s.stop(false)
			logx.Errorf("grpcx: %v", err)
			return err
		}

		return nil
	}
}

// ServeAndWait - start the server and wait for its stopping.
// See document of Server.Serve for more information.
func (s *Server) ServeAndWait() error {
	if err := s.Serve(); err != nil {
		return err
	} else {
		s.waitGroup.Wait()
		return nil
	}
}

// Stop - this function stops the server by call to grpc.Server.GracefulStop.
// The BeforeStopHook will be called before the method is called.
func (s *Server) Stop() {
	s.stop(true)
}

func (s *Server) stop(invokeHook bool) {
	if s.isStart {
		// executes all BeforeStopHook hooks.
		if invokeHook {
			s.execBeforeStopHook()
		}

		s.realServer.GracefulStop()
		s.isStart = false
	}
}

// start - starts the server.
func (s *Server) start() error {
	s.execRegisterRpcHook()

	listener, err := NewListenerWithCfg(s.cfg)
	if err != nil {
		return err
	}

	s.waitGroup.Add(1)
	go func() {
		s.isStart = true

		logx.Infof("grpcx: server is serve on %v:%v", s.cfg.Addr, s.cfg.Port)
		if err = s.realServer.Serve(listener); err != nil {
			logx.Fatal("grpcx: ", err)
			return
		}

		s.isStart = false
		s.waitGroup.Done()
	}()

	return nil
}

func (s *Server) execRegisterRpcHook() {
	for _, server := range s.servers {
		server.RegisterRpc(s.realServer)
	}
}

func (s *Server) execBeforeStartHook() error {
	for _, server := range s.servers {
		if hook, ok := server.(BeforeStartHook); ok {
			if err := hook.BeforeServerStart(s.cfg, s); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Server) execAfterStartHook() error {
	for _, server := range s.servers {
		if hook, ok := server.(AfterStartHook); ok {
			if err := hook.AfterServerStart(s.cfg, s); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Server) execBeforeStopHook() {
	for _, server := range s.servers {
		if hook, ok := server.(BeforeStopHook); ok {
			hook.BeforeServerStop(s.cfg, s)
		}
	}
}

// NewServer - creates a grpc server with given options.
func NewServer(opts ...OptionFunc) *Server {
	server := &Server{realServer: grpc.NewServer()}
	for _, apply := range opts {
		apply(server)
	}

	if server.cfg == nil {
		logx.Info("grpcx: use default configuration")
		server.cfg = gCfg
	}

	return server
}
