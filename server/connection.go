package server

import (
	"fmt"
	"log"
	"net"
)

func (server *Server) Accept() {
	for {
		conn, err := server.Listener.Accept()
		if err != nil {
			log.Fatal("error accepting connection:", err)
		}

		server.handleRequest(conn)
	}
}

func (server *Server) handleRequest(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, MAX_BUFFER_SIZE)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("error reading connection:", err)
		return
	}

	httpProtocol := server.parseRequest(buffer, n)
	fmt.Printf("%v\n%v\n", httpProtocol, *httpProtocol.URL)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
}
