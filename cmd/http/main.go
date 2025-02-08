package main

import (
	"fmt"

	"github.com/ashiqYousuf/http-go-server/server"
)

func main() {
	server := server.NewServer(8000, "127.0.0.1")

	fmt.Println("server listening on ", fmt.Sprintf("%s:%d", "127.0.0.1", 8000))

	server.Accept()
}
