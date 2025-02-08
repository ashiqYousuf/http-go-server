# HTTP Server in Golang from Scratch

## Introduction
This project is a simple HTTP server implemented in Go, built from scratch for learning purposes. It helps you understand how HTTP requests and responses work at a fundamental level without using any frameworks.

---

## Features
- Handles concurrent HTTP requests
- Demonstrates request parsing
- Responds proper HTTP response including Headers and Status Codes
- Lightweight and easy to modify

---

## HTTP Request & Response Format
HTTP communication follows a structured format:

```
<Request Line or Status Line>   --> e.g., "GET /path HTTP/1.1" \r\n
<Headers>                      --> e.g., "Host: 127.0.0.1" (headers are separated by \r\n)
(blank line)                   --> "\r\n" (End of headers)
<Body (optional)>              --> (Present in POST, PUT, etc.)
```

### Example HTTP Request (from client to server)
```
GET / HTTP/1.1

Host: 127.0.0.1:8080

User-Agent: curl/7.68.0

Accept: */*



```

### Example HTTP Response (from server to client)
```
HTTP/1.1 200 OK

Content-Type: text/plain

Content-Length: 12



Hello World!
```

---

## Running the Server
1. Install Go if not already installed: [Download Go](https://go.dev/dl/)
2. Clone this repository:
   ```sh
   git clone https://github.com/your-username/http-go-server.git
   cd http-go-server
   ```
3. Run the server:
   ```sh
   go run main.go
   ```
4. Open another terminal and send a request:
   ```sh
   curl -v http://127.0.0.1:8080
   ```

---


## Contributing
Feel free to contribute by submitting a pull request or opening an issue. This project serves a great learning purpose.
