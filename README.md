# Halo

A lightweight and flexible Go web framework

## Installation

```sh
go get github.com/yinfxs/halo
```

## Import

```go
import "github.com/yinfxs/halo"
```

## Usage

```go
package main

import (
  "fmt"
  "log"
  "time"

  "github.com/yinfxs/halo"
)

// Logger 日志中间件
func Logger(ctx *halo.Context) {
  start := time.Now().Unix()
  raw := ctx.R.URL.RawQuery

  info := fmt.Sprintf("%s %s", ctx.R.Method, ctx.R.URL.Path)
  if raw != "" {
    info += "?" + raw
  }

  log.Printf("--> %v\n", info)
  ctx.Next()
  log.Printf("<-- %s %dms\n", info, start-time.Now().Unix())
}

func main() {
  a := halo.New()
  a.Use(Logger)
  a.Run(":3000")
}
```

```sh
2018/06/27 18:47:30 Listen and serve on :3000
2018/06/27 18:47:33 --> GET /
2018/06/27 18:47:33 <-- GET / 0ms
2018/06/27 18:47:34 --> GET /favicon.ico
2018/06/27 18:47:34 <-- GET /favicon.ico 0ms
```

## Contributing

If you'd like to help out with the project. You can put up a Pull Request.
