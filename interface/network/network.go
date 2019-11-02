package network

import (
	"context"
	"io"
)

type ApiResponser interface {
	GetResponseWriter() ResponseWriter
	GetRequest() Request
	SetRequestContext(context.Context)
	GetRequestContext() context.Context
	Success(interface{})
	BadRequset(string)
	InternalServerError(string)
}

type ResponseWriter interface {
	Write([]byte) (int, error)
	WriteHeader(int)
}
type Request interface {
	GetBody() io.Reader
	GetDeaderValue(string) string
}