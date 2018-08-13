package request

import (
	"bufio"
	"io"
)

type Request struct {
	RequestLine
	Headers map[string]string
	Body    string
}

func ReadRequest(buf io.Reader) *Request {
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanLines)

	// request-line
	scanner.Scan()
	requestLine := readRequestLine(scanner.Bytes())

	// headers
	headers := make(map[string]string)
	scanner.Scan()
	l := scanner.Bytes()
	for len(l) > 0 {
		fieldName, fieldContent := readHeader(l)
		headers[fieldName] = fieldContent
		scanner.Scan()
		l = scanner.Bytes()
	}

	return &Request{
		*requestLine,
		headers,
		"",
	}
}
