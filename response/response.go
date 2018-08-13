package response

import (
	"bytes"
	"fmt"
)

type Response struct {
	StatusLine
	Headers map[string]string
	Body    string
}

func NewResponse(statusCode int, headers map[string]string, body string) *Response {
	statusLine := newStatusLine(statusCode)
	if statusLine == nil {
		panic(fmt.Sprintf("%d is not a valid status code", statusCode))
	}

	return &Response{
		*statusLine,
		headers,
		body,
	}
}

func (res *Response) Bytes() []byte {
	// TODO: add headers to serializer
	values := [][]byte{
		res.StatusLine.Bytes(),
		[]byte(""),
		[]byte(res.Body),
	}
	return bytes.Join(values, []byte("\r\n"))
}
