package main

import (
	"flag"
	"log"
	"os"

	"github.com/ashiqYousuf/http-go-server/internal/server"
	"github.com/ashiqYousuf/http-go-server/pkg/clogger"
)

var (
	PORT = 8000
	HOST = "127.0.0.1"
)

func main() {
	port := flag.Int("port", PORT, "port to listen on")
	host := flag.String("host", HOST, "network host")

	errorLogger := clogger.ErrorLogger(os.Stderr, log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger := clogger.InfoLogger(os.Stdout, log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger := clogger.WarningLogger(os.Stdout, log.Ldate|log.Ltime|log.Lshortfile)
	successLogger := clogger.SuccessLogger(os.Stdout, log.Ldate|log.Ltime|log.Lshortfile)

	flag.Parse()

	server := server.NewServer(*port, *host, infoLogger, errorLogger, warningLogger, successLogger)
	server.ListenAndServe()
}
