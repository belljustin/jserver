package main

import (
	"bytes"
	"fmt"
	"io"
	"net"

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
	var buf bytes.Buffer

	tmp := make([]byte, 256)
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				// TODO: handle error
			}
			break
		}

		buf.Write(tmp[:n])

		// TODO: replace with reading length
		if tmp[n-2] == '\r' && tmp[n-1] == '\n' {
			break
		}
	}

	req := request.ReadRequest(&buf)
	fmt.Printf("%+v\n", req)
	// TODO: actually do something with request

	resHeaders := make(map[string]string)
	res := response.NewResponse(200, resHeaders, "hello")
	conn.Write(res.Bytes())
}
