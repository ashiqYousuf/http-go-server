package server

import (
	"fmt"
	"log"
	"net"
)

type HTTPServer struct {
	Addr        string
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
	Listener    *net.Listener
}

func NewServer(port int, host string, infoLogger *log.Logger, errorLogger *log.Logger) *HTTPServer {
	addr := fmt.Sprintf("%s:%d", host, port)

	return &HTTPServer{
		Addr:        addr,
		ErrorLogger: errorLogger,
		InfoLogger:  infoLogger,
	}
}
