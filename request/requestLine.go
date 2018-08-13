package request

import (
	"fmt"
	"regexp"

	"github.com/belljustin/jserver/hashset"
)

type RequestLine struct {
	Method     string
	RequestURI string
	Version    string
}

func readRequestLine(b []byte) *RequestLine {
	matches := requestLineRegex.FindSubmatch(b)
	// TODO: error checking

	method := string(matches[1])
	requestURI := string(matches[2])
	version := string(matches[3])

	return &RequestLine{
		method,
		requestURI,
		version,
	}
}

// requestLineRegex matches "method SP requestURI SP version"
var requestLineRegex = func() *regexp.Regexp {
	r := fmt.Sprintf("^(%s) (%s) (%s)$", methodsRegex, requestUriRegex,
		versionRegex)
	return regexp.MustCompile(r)
}()

// Methods
var Methods = hashset.NewStringSet(
	"OPTIONS",
	"GET",
	"HEAD",
	"POST",
	"PUT",
	"DELETE",
	"TRACE",
	"CONNECT",
)

var methodsRegex = func() string {
	re := regexp.MustCompile(",")
	return re.ReplaceAllString(Methods.String(), "|")
}()

// Request URI
// TODO: use a more sophisticated uri regex
const requestUriRegex = ".*"

// Version
const versionRegex = "HTTP/[1-9]+.[1-9]+"
