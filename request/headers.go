package request

import (
	"regexp"

	"github.com/belljustin/jserver/hashset"
)

var HEADERS = hashset.NewStringSet(
	// general-header
	"Cache-Control",
	"Connection",
	"Date",
	"Pragma",
	"Trailer",
	"Transfer-Encoding",
	"Upgrade",
	"Via",
	"Warning",
	// request-header
	"Accept",
	"Accept-Charset",
	"Accept-Encoding",
	"Accept-Language",
	"Authorization",
	"Expect",
	"From",
	"Host",
	"If-Match",
	"If-Modified-Since",
	"If-None-Match",
	"If-Range",
	"If-Unmodified-Since",
	"Max-Forwards",
	"Proxy-Authorization",
	"Range",
	"Referer",
	"TE",
	"User-Agent",
	// entity-header
	"Allow",
	"Content-Encoding",
	"Content-Language",
	"Content-Length",
	"Content-Location",
	"Content-MD5",
	"Content-Range",
	"Content-Type",
	"Expires",
	"Last-Modified",
	"extension-header",
)

func readHeader(b []byte) (string, string) {
	r := regexp.MustCompile("^([^:]*):(?: )*?([^ ].*)$")
	loc := r.FindSubmatchIndex(b)
	// TODO check if valid
	fieldName := string(b[loc[2]:loc[3]])
	fieldContent := string(b[loc[4]:loc[5]])
	return fieldName, fieldContent
}
