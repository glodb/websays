package baserouter

import (
	"sync"

	"github.com/gorilla/mux"
)

// baseRouterHandler is a singleton object responsible for managing named routers in the application.
type baseRouterHandler struct {
	router map[string]*mux.Router // A map to store named routers.
}

var instance *baseRouterHandler // Singleton instance.
var once sync.Once              // Ensures that instance creation is thread-safe.

// GetInstance returns a single object of baseRouterHandler.
func GetInstance() *baseRouterHandler {
	// Create a single instance if it doesn't exist.
	once.Do(func() {
		instance = &baseRouterHandler{}
		instance.router = make(map[string]*mux.Router)
	})
	return instance
}

// SetRouter associates a named router with a specific name.
func (u *baseRouterHandler) SetRouter(name string, router *mux.Router) {
	u.router[name] = router // Store the named router in the map.
}

// GetBaseRouter retrieves the "base" named router.
// It allows access to the main router for adding routes or middleware.
func (u *baseRouterHandler) GetBaseRouter() *mux.Router {
	return u.router["base"] // Retrieve and return the "base" named router.
}
