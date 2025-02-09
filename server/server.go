package server

import (
	"fmt"
	"log"
	"net"
)

var (
	MAX_BUFFER_SIZE = 1024
)

type Server struct {
	Port     int
	Listener net.Listener
	Routes   []*Route
}

func NewServer(port int, host string) *Server {
	addr := fmt.Sprintf("%s:%d", host, port)
	fmt.Println(addr)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("failed to bind to port", port)
	}
	return &Server{
		Port:     port,
		Listener: listener,
		Routes:   make([]*Route, 0),
	}
}

func (server *Server) AddRoute(method string, path string, cb func(*HTTPRequest) *HTTPResponse) {
	server.Routes = append(server.Routes, &Route{
		Callback: cb,
		Method:   method,
		Path:     path,
	})
}
