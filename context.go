package halo

import (
	"net/http"

	"github.com/yinfxs/middleware"
)

// Context 上下文对象
type Context struct {
	context  *middleware.Context
	Next     func()
	Data     M
	W        http.ResponseWriter
	R        *http.Request
	Request  Request
	Response Response
	Params   M
}
