package clogger

import (
	"io"
	"log"
)

func NewLogger(w io.Writer, prefix string, flag int) *log.Logger {
	return log.New(w, prefix, flag)
}

func InfoLogger(w io.Writer, flag int) *log.Logger {
	return NewLogger(w, Cyan, flag)
}

func ErrorLogger(w io.Writer, flag int) *log.Logger {
	return NewLogger(w, Red, flag)
}
func WarningLogger(w io.Writer, flag int) *log.Logger {
	return NewLogger(w, Yellow, flag)
}

func SuccessLogger(w io.Writer, flag int) *log.Logger {
	return NewLogger(w, Green, flag)
}
