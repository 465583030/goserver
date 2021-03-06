Vardius - goserver
================
[![Build Status](https://travis-ci.org/vardius/goserver.svg?branch=master)](https://travis-ci.org/vardius/goserver) [![](https://godoc.org/github.com/vardius/goserver?status.svg)](http://godoc.org/github.com/vardius/goserver) [![Coverage Status](https://coveralls.io/repos/github/vardius/goserver/badge.svg?branch=master)](https://coveralls.io/github/vardius/goserver?branch=master)

Go Server/API micro framwework, HTTP request router, multiplexer, mux.

Appling Middleware
----------------
1. [Global Middlewares](#global-middlewares)
2. [Method Middlewares](#method-middlewares)
3. [Route Middlewares](#route-middlewares)

## Global Middlewares
```go
package main

import (
	"fmt"
	"log"
	"time"
	"net/http"

    "github.com/vardius/goserver"
)

func logger(next http.Handler) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
    t1 := time.Now()
    next.ServeHTTP(w, r)
    t2 := time.Now()
    log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
  }

  return http.HandlerFunc(fn)
}

func example(next http.Handler) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
	//do smth
    next.ServeHTTP(w, r)
  }

  return http.HandlerFunc(fn)
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	params, _ := goserver.ParamsFromContext(r.Context())
    fmt.Fprintf(w, "hello, %s!\n", params["name"])
}

func main() {
	//apply middlewares to all routes
	//can pass as many as you want
    server := goserver.New(logger, example)

    server.GET("/", Index)
    server.GET("/hello/:name", Hello)

    log.Fatal(http.ListenAndServe(":8080", server))
}
```
## Method Middlewares
```go
package main

import (
	"fmt"
	"log"
	"time"
	"net/http"

    "github.com/vardius/goserver"
)

func logger(next http.Handler) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
    t1 := time.Now()
    next.ServeHTTP(w, r)
    t2 := time.Now()
    log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
  }

  return http.HandlerFunc(fn)
}

func example(next http.Handler) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
	//do smth
    next.ServeHTTP(w, r)
  }

  return http.HandlerFunc(fn)
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	params, _ := goserver.ParamsFromContext(r.Context())
    fmt.Fprintf(w, "hello, %s!\n", params["name"])
}

func main() {
    server := goserver.New()

    server.GET("/", Index)
    server.GET("/hello/:name", Hello)

	//apply middlewares to all routes with GET method
	//can pass as many as you want
    server.USE("GET", "", logger, example)

    log.Fatal(http.ListenAndServe(":8080", server))
}
```
## Route Middlewares
```go
package main

import (
	"fmt"
	"log"
	"time"
	"net/http"

    "github.com/vardius/goserver"
)

func logger(next http.Handler) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
    t1 := time.Now()
    next.ServeHTTP(w, r)
    t2 := time.Now()
    log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
  }

  return http.HandlerFunc(fn)
}

func example(next http.Handler) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
	//do smth
    next.ServeHTTP(w, r)
  }

  return http.HandlerFunc(fn)
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	params, _ := goserver.ParamsFromContext(r.Context())
    fmt.Fprintf(w, "hello, %s!\n", params["name"])
}

func main() {
    server := goserver.New()
    server.GET("/", Index)

    server.GET("/hello/:name", Hello)

	//apply midlewares to route and all it childs
	//can pass as many as you want
    server.USE("GET", "/hello/:name", logger, example)

    log.Fatal(http.ListenAndServe(":8080", server))
}
```
