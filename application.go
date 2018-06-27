package halo

import (
	"log"
	"net/http"
	"sync"

	"github.com/yinfxs/middleware"
)

// M is a shortcup for map[string]interface{}
type M map[string]interface{}

// App 应用
type App struct {
	addr string
	mw   *middleware.Application
	pool sync.Pool
	c    *Context
}

// Use 应用
func (a *App) Use(fn func(ctx *Context)) {
	a.mw.Add(func(ctx *middleware.Context) {
		fn(a.c)
	})
}

// output 输出监听日志
func (a *App) output() {
	if a.addr != "" {
		log.Printf("Listen and serve on %v", a.addr)
	}
}

// Run 启动应用
func (a *App) Run(addr string) (err error) {
	a.addr = addr
	a.output()
	if err := http.ListenAndServe(addr, a); err != nil {
		panic(err)
	}
	return nil
}

// RunTLS 启动应用（TLS方式）
func (a *App) RunTLS(addr, certFile, keyFile string) (err error) {
	a.addr = addr
	a.output()
	if err := http.ListenAndServeTLS(addr, certFile, keyFile, a); err != nil {
		panic(err)
	}
	return nil
}

// AcquireContext 请求上下文对象
func (a *App) AcquireContext() *Context {
	return a.pool.Get().(*Context)
}

// ReleaseContext 释放上下文对象回对象池
func (a *App) ReleaseContext() {
	a.pool.Put(a.c)
}

// NewContext 创建上下文对象
func (a *App) NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Data: M{},
		R:    r,
		W:    w,
	}
}

// ServeHTTP 实现接口 http.Handler
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mw.Flow(func(context *middleware.Context) {
		c := a.AcquireContext()
		c.context = context
		c.Next = context.Next
		c.R = r
		c.W = w
		a.c = c
	})
	a.ReleaseContext()
}

// New 创建应用实例
func New() *App {
	a := &App{
		mw: middleware.New(),
	}
	a.pool.New = func() interface{} {
		return a.NewContext(nil, nil)
	}
	return a
}
