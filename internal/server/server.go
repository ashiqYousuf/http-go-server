package server

import (
	"fmt"
	"log"
	"net"
)

type Server interface {
	ListenAndServe()
}

type HTTPServer struct {
	Addr          string
	Listener      *net.Listener
	ErrorLogger   *log.Logger
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	SuccessLogger *log.Logger
}

func NewServer(port int, host string, infoLogger *log.Logger, errorLogger *log.Logger, warningLogger *log.Logger, successLogger *log.Logger) *HTTPServer {
	addr := fmt.Sprintf("%s:%d", host, port)

	return &HTTPServer{
		Addr:          addr,
		ErrorLogger:   errorLogger,
		InfoLogger:    infoLogger,
		WarningLogger: warningLogger,
		SuccessLogger: successLogger,
	}
}
