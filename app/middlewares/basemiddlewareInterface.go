package middlewares

import "net/http"

// Middleware is the base middleware interface used for writing middleware components.
// A middleware typically enhances or modifies HTTP request handling by adding specific
// functionality. Implementations of this interface should provide a GetHandlerFunc method
// to define the middleware's handler function, which will be used to process incoming
// HTTP requests.
type Middleware interface {
	// GetHandlerFunc returns an HTTP handler function that wraps the provided 'next'
	// HTTP handler. This allows the middleware to intercept and process HTTP requests
	// before they reach the 'next' handler in the chain.
	GetHandlerFunc(next http.Handler) http.Handler
}
