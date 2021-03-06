package goserver

import "net/http"

type (
	MiddlewareFunc func(http.Handler) http.Handler
	middleware     []MiddlewareFunc
)

func (m middleware) handle(h http.Handler) http.Handler {
	if h == nil {
		h = http.DefaultServeMux
	}

	for i := range m {
		h = m[len(m)-1-i](h)
	}

	return h
}

func (m middleware) handleFunc(f http.HandlerFunc) http.Handler {
	return m.handle(f)
}

func (m middleware) append(fs ...MiddlewareFunc) middleware {
	return append(m, fs...)
}

func (m middleware) merge(n middleware) middleware {
	return append(m, n...)
}

func newMiddleware(fs ...MiddlewareFunc) middleware {
	ms := make(middleware, 0, len(fs))
	for _, f := range fs {
		if f != nil {
			ms = append(ms, f)
		}
	}

	return ms
}
