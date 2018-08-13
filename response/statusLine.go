package response

import (
	"bytes"
	"strconv"
)

const version = "HTTP/1.1"

type StatusLine struct {
	Version      string
	StatusCode   int
	ReasonPhrase string
}

func newStatusLine(statusCode int) *StatusLine {
	reasonPhrase, found := StatusCodes[statusCode]
	if found == false {
		return nil
	}

	return &StatusLine{
		version,
		statusCode,
		reasonPhrase,
	}
}

func (sl *StatusLine) Bytes() []byte {
	values := [][]byte{
		[]byte(sl.Version),
		[]byte(strconv.Itoa(sl.StatusCode)),
		[]byte(sl.ReasonPhrase),
	}
	return bytes.Join(values, []byte(" "))
}

var StatusCodes = map[int]string{
	200: "Ok",
}
