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
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Handling connection")

	buf, bodyBuf := readRequestLineAndHeaders(conn)
	req := request.ReadRequest(buf, bodyBuf)
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

	resHeaders := make(map[string]string)
	res := response.NewResponse(200, resHeaders, bodyBuf.String())
	conn.Write(res.Bytes())
}

func readRequestLineAndHeaders(r io.Reader) (buf bytes.Buffer, bodyBuf bytes.Buffer) {
	tmp := make([]byte, 256)

	// delim CRLF CRLF marks end of request line and headers
	delim := [...]byte{'\r', '\n', '\r', '\n'}
	i := 0

	for {
		n, err := r.Read(tmp)
		if err != nil {
			if err != io.EOF {
				panic("Encountered an error parsing request line an headers")
			}
		}

		for m, b := range tmp {
			// increment the counter if the next delim is found
			if b == delim[i] {
				i += 1
			} else {
				i = 0
			}

			// if found entire delim, write part before to buf and the rest to
			// the body buffer
			if i == len(delim) {
				buf.Write(tmp[:m])
				bodyBuf.Write(tmp[m:n])
				return buf, bodyBuf
			}
		}

		buf.Write(tmp[:n])
	}
	return buf, bodyBuf
}
