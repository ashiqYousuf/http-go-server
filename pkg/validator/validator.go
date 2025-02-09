package validator

func BufferExceeded(bytesRead, bytesLimit int) bool {
	return bytesRead > bytesLimit
}
