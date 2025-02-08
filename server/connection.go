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

		go server.handleRequest(conn)
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

	httpRequest := server.parseRequest(buffer, n)
	httpResponse := server.routeRequests(httpRequest)
	fmt.Printf("%v\n", httpRequest.URL)

	bytesBuffer, err := server.convertHTTPProtocolToBytes(&httpResponse.HTTPProtocol)
	if err != nil {
		fmt.Println("error writing response:", err)
		return
	}

	conn.Write(bytesBuffer)
}
