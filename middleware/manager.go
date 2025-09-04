package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next

	/// For single middleware
	for _, middleware := range middlewares {
		n = middleware(n)
	}
	return n
}

func (mngr *Manager) WrapMux(next http.Handler) http.Handler {
	n := next

	for i := len(mngr.globalMiddlewares) - 1; i >= 0; i-- {
		// Apply each middleware in reverse order
		n = mngr.globalMiddlewares[i](n)
	}

	// for _, globalMiddleWares := range mngr.globalMiddlewares {
	// 	n = globalMiddleWares(n)
	// }
	return n
}
