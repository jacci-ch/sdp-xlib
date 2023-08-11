// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package grpcx

import (
	"errors"
	"github.com/jacci-ch/sdp-xlib/logx"
	"google.golang.org/grpc"
	"sync"
)

type Server struct {
	Cfg     *Config
	isStart bool

	waitGroup  sync.WaitGroup
	realServer *grpc.Server

	userServer GrpcServer
}

// Serve
//
// Start the server. The rpc.register and hooks will be called
// at the appropriate time.
func (s *Server) Serve() error {
	if s.isStart {
		return errors.New("grpcx: server is already started")
	}

	if hook, ok := s.userServer.(BeforeStartHook); ok {
		if err := hook.BeforeServerStart(s); err != nil {
			return err
		}
	}

	if err := s.start(); err != nil {
		return err
	}

	if hook, ok := s.userServer.(AfterStartHook); ok {
		if err := hook.AfterServerStart(s); err != nil {
			s.Stop()
			return err
		}
	}

	return nil
}

// ServeAndWait
//
// Start the server and wait for its stopping.
func (s *Server) ServeAndWait() error {
	if err := s.Serve(); err != nil {
		return err
	}

	s.waitGroup.Wait()
	return nil
}

// start
//
// Starts the server.
func (s *Server) start() error {
	if err := s.userServer.Register(s.realServer); err != nil {
		return err
	}

	listener, err := NewListenerWithCfg(s.Cfg)
	if err != nil {
		return err
	}

	s.waitGroup.Add(1)
	go func() {
		s.isStart = true
		logx.Logger.Infof("rgpcx: server is serve on %v", s.Cfg.Endpoint())
		if err = s.realServer.Serve(listener); err != nil {
			logx.Logger.Fatal("grpcx: ", err)
		}
		s.isStart = false
		s.waitGroup.Done()
	}()

	return nil
}

// Stop
//
// This function stop the server by call to grpc.Server.GracefulStop function.
// The BeforeStopHook will be called before the method is called.
func (s *Server) Stop() {
	if !s.isStart {
		return
	}

	if hook, ok := s.userServer.(BeforeStopHook); ok {
		hook.BeforeServerStop(s)
	}

	s.realServer.GracefulStop()
	s.isStart = false
}

// NewServerWithCfg
//
// New a server with given configurations and the register.
func NewServerWithCfg(cfg *Config, server GrpcServer) (*Server, error) {
	if cfg == nil {
		return nil, errors.New("grpcx: argument cfg can't be nil")
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	s := grpc.NewServer()
	return &Server{Cfg: cfg, realServer: s, userServer: server}, nil
}

// NewServerWithKeys
//
// New a server with given configuration keys and the register. This function
// loads all configurations and then call to NewServerWithCfg to create a
// new server object.
func NewServerWithKeys(keys *ConfigKeys, server GrpcServer) (*Server, error) {
	cfg, err := LoadConfigsWith(keys)
	if err != nil {
		return nil, err
	}

	return NewServerWithCfg(cfg, server)
}

// NewServer
//
// New a server with default configurations and the given register.
func NewServer(server GrpcServer) (*Server, error) {
	return NewServerWithCfg(Cfg, server)
}
