package baserouter

import (
	"github.com/gorilla/mux"
)

// BaseRouter is an interface that defines methods for managing HTTP routers in an application.
type BaseRouter interface {
	// SetRouter sets the router with a given name.
	// It associates a named router instance with a specific purpose or route group.
	SetRouter(name string, router *mux.Router)

	// GetOpenRouter retrieves the router associated with the given name.
	// It allows access to the named router, typically for adding routes or middleware.
	GetRouter(name string) *mux.Router
}
