package server

import (
	"fmt"
	"net"
	"strconv"

	"github.com/ashiqYousuf/http-go-server/internal/protocol"
	"github.com/ashiqYousuf/http-go-server/internal/status"
	"github.com/ashiqYousuf/http-go-server/pkg/utils"
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

	httpRequest, err := protocol.ParseRequest(conn)
	if err != nil {
		server.ErrorLogger.Println("error reading connection, err:", err)
	}

	fmt.Printf("REQUEST:\n%v\n%v\n", httpRequest.URL, httpRequest.HTTPProtocol)

	response := "Hello, World!"
	httpProtocol := protocol.NewHTTPProtocol(map[string]string{
		"Content-Type":   "text/plain",
		"Content-Length": strconv.Itoa(len(response)),
	}, response)

	httpResponse := protocol.NewHTTPResponse(httpProtocol, status.StatusOK)
	responseBytes, _ := protocol.HTTPResponseToBytes(httpResponse)
	utils.WriteBytes(conn, responseBytes)

	// TODO:- parse requests
	// Do routing
	// Server HTTP Responses
}
