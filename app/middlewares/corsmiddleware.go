package middlewares

import (
	"errors"
	"net/http"
	"websays/httpHandler/responses"
)

// CORSMiddleware is a middleware for handling Cross-Origin Resource Sharing (CORS) in HTTP requests.
// CORS middleware allows or restricts web applications running at one origin (domain) to make
// requests to a different origin (domain). It adds the necessary CORS headers to responses to
// enable cross-origin requests.
type CORSMiddleware struct {
}

// GetHandlerFunc returns an HTTP handler function for CORS middleware.
// It sets the necessary CORS headers to allow cross-origin requests and handles OPTIONS requests
// for pre-flight checks.
//
// Parameters:
//   - next: The next HTTP handler in the middleware chain.
//
// Returns:
//   - http.Handler: An HTTP handler function that wraps the provided 'next' handler and adds CORS
//     headers to the response.
func (c *CORSMiddleware) GetHandlerFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers to allow cross-origin requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, mode, Access-Control-Allow-Origin, x-access-token, ssoSession, url")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		// Handle OPTIONS request for pre-flight checks
		if r.Method == "OPTIONS" {
			// Respond with an error message for OPTIONS requests
			responses.GetInstance().WriteJsonResponse(w, r, responses.OPTIONS_NOT_ALLOWED, errors.New("Method Options is not allowed in Cors Middleware"), nil)
			return
		}

		// Call the next handler in the middleware chain
		next.ServeHTTP(w, r)
	})
}
