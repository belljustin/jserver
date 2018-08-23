# jserver
An http server written in Go - a learning excersise.

## Installation & Running

```sh
go get 
go build

./jserver
```

### Testing

```sh
$ curl http://localhost:9090 -X POST -d "hello world"
```

## Relevant Documentation

### RFCs
- [rfc7230](https://tools.ietf.org/html/rfc7230) - Hypertext Transfer Protocol (HTTP/1.1): Message Syntax and Routing
- [rfc7231](https://tools.ietf.org/html/rfc7231) - Hypertext Transfer Protocol (HTTP/1.1): Semantics and Content

- [rfc2616](https://tools.ietf.org/html/rfc2616) - Hypertext Transfer Protocol -- HTTP/1.1
    - Obseleted by rfc7230 including removing OWS for line folding on headers
- [rfc2396](https://tools.ietf.org/html/rfc2396) - Uniform Resource Identifiers (URI): Generic Syntax)
- [rfc822](https://tools.ietf.org/html/rfc822) - STANDARD FOR THE FORMAT OF ARPA INTERNET TEXT MESSAGES
