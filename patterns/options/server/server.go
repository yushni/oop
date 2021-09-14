package server

import (
	"fmt"
	"time"
)

type Opts func(*Server)

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

func New(options ...Opts) *Server {
	svr := &Server{}
	for _, o := range options {
		o(svr)
	}
	return svr
}

func (s *Server) Start() error {
	fmt.Println("ми почали почали", s)
	return nil
}

func WithHost(host string) Opts {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) Opts {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) Opts {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConn(maxConn int) Opts {
	return func(s *Server) {
		s.maxConn = maxConn
	}
}
