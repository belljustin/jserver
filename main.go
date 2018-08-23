package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"strconv"

	"github.com/belljustin/jserver/request"
	"github.com/belljustin/jserver/response"
)

const (
	HOST = "localhost"
	PORT = 9090
)

func main() {
	listen := fmt.Sprintf("%s:%d", HOST, PORT)

	ln, err := net.Listen("tcp", listen)
	if err != nil {
		// TODO: handle error
		fmt.Println(err)
	}

	fmt.Printf("Accepting connections on %s\n", listen)
	for {
		conn, err := ln.Accept()
		if err != nil {
			// TODO: handle error
			fmt.Println(err)
		}
		go handleConnection(conn, echoHandler)
	}
}

func handleConnection(conn net.Conn, handle handler) {
	defer conn.Close()

	fmt.Println("Handling connection")

	buf, bodyBuf := readHeader(conn)
	req := request.ParseRequest(buf, bodyBuf)
	fmt.Printf("%+v\n", req)

	// If request is POST, use Content-Length header to read rest of request
	if req.Method == "POST" {
		n, found := req.Headers["Content-Length"]
		if !found {
			panic("No Content-Length for POST request")
		}

		m, _ := strconv.Atoi(n) // TODO: handle error
		tmp := make([]byte, (m+1)-bodyBuf.Len())
		conn.Read(tmp) // TODO: handle error
		bodyBuf.Write(tmp)
	}

	res := handle(req)
	conn.Write(res.Bytes())
}

func readHeader(r io.Reader) (headerBuf bytes.Buffer, bodyBuf bytes.Buffer) {
	tmp := make([]byte, 256)

	// delim CRLF CRLF marks end of request line and headers
	delim := [...]byte{'\r', '\n', '\r', '\n'}
	i := 0

	for {
		n, _ := r.Read(tmp)

		for m, b := range tmp {
			// increment the counter if the next delim is found
			if b == delim[i] {
				i += 1
			} else {
				i = 0
			}

			// if entire delim is found, write part before to headerBuf and
			// the rest to the body buffer
			if i == len(delim) {
				headerBuf.Write(tmp[:m])
				bodyBuf.Write(tmp[m:n])
				return headerBuf, bodyBuf
			}
		}

		headerBuf.Write(tmp[:n])
	}
	return headerBuf, bodyBuf
}

// Hanlders

type handler func(*request.Request) *response.Response

func echoHandler(req *request.Request) *response.Response {
	headers := make(map[string]string)
	return response.NewResponse(200, headers, req.Body)
}
