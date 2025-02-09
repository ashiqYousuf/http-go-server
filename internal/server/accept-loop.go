package server

import (
	"net"

	"github.com/ashiqYousuf/http-go-server/internal/constants"
	"github.com/ashiqYousuf/http-go-server/pkg/validator"
)

func (server *HTTPServer) acceptLoop() {
	for {
		conn, err := (*server.Listener).Accept()
		if err != nil {
			server.ErrorLogger.Fatal("error accepting connection, err:", err)
		}

		go server.handleConnection(conn)
	}
}

func (server *HTTPServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, constants.MAX_BUFFER_SIZE)
	n, err := conn.Read(buffer)
	if err != nil {
		server.ErrorLogger.Println("error reading connection, err:", err)
		// TODO:
		return
	}

	if validator.BufferExceeded(n, constants.MAX_BUFFER_SIZE) {
		server.ErrorLogger.Println("error reading connection, err:", err)
		// TODO:
		return
	}

	// 	httpRequest := server.parseRequest(buffer, n)
	// 	httpResponse := server.routeRequests(httpRequest)
	// 	fmt.Printf("%v\n", httpRequest.URL)

	// 	bytesBuffer, err := server.convertHTTPProtocolToBytes(&httpResponse.HTTPProtocol)
	// 	if err != nil {
	// 		fmt.Println("error writing response:", err)
	// 		return
	// 	}

	// 	conn.Write(bytesBuffer)
	// }

}
