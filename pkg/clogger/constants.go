package clogger

// ANSI escape codes
const (
	Reset  = "\033[0m"
	Red    = "\033[31m" + "[ERROR] "
	Yellow = "\033[33m" + "[WARNING] "
	Green  = "\033[32m" + "[SUCCESS] "
	Cyan   = "\033[36m" + "[INFO] "
)
