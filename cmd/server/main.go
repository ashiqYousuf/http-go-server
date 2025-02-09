package main

import (
	"flag"
	"log"
	"os"

	"github.com/ashiqYousuf/http-go-server/internal/server"
)

var (
	PORT = 8000
	HOST = "127.0.0.1"
)

func main() {
	port := flag.Int("port", PORT, "port to listen on")
	host := flag.String("host", HOST, "network host")

	errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)

	flag.Parse()

	server := server.NewServer(*port, *host, infoLogger, errorLogger)

	server.ListenAndServe()
}
