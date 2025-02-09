package server

import "net"

func (server *HTTPServer) ListenAndServe() {
	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		server.ErrorLogger.Fatal("failed to bind to address: ", server.Addr, "err: ", err)
	}
	defer listener.Close()

	server.Listener = &listener
	server.acceptLoop()
}
