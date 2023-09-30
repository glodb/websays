package httpHandler

import (
	"log"
	"net/http"
	"sync"
	"websays/app/middlewares"
	"websays/config"
	"websays/httpHandler/basecontrollers"
	"websays/httpHandler/baserouter"
	"websays/httpHandler/responses"

	"github.com/gorilla/mux"
)

// muxServer is a struct that represents the Mux server.
type muxServer struct {
	base *mux.Router // Mux router instance.
}

var (
	instance *muxServer // Singleton instance of muxServer.
	once     sync.Once  // Used for ensuring singleton behavior.
)

// GetInstance returns a singleton instance of the muxServer.
// It sets up the instance and its associated routes.
func GetInstance() *muxServer {
	once.Do(func() {
		instance = &muxServer{}
		instance.setup()
	})
	return instance
}

// HandleBlank is a handler function for a blank route ("/").
// It uses the responses package to write a JSON response.
func (u *muxServer) HandleBlank(w http.ResponseWriter, r *http.Request) {
	// Use the responses package to write a JSON response.
	responses.GetInstance().WriteJsonResponse(w, r, responses.WEBSAYS_TEST, nil, nil)
}

// setup configures the Mux router and sets up necessary middleware.
func (u *muxServer) setup() {
	corsMiddleware := middlewares.CORSMiddleware{} // Initialize CORS middleware.
	u.base = &mux.Router{}                         // Initialize the Mux router.

	// Use CORS middleware for all routes handled by this router.
	u.base.Use(corsMiddleware.GetHandlerFunc)

	// Define a route for the root path ("/") and associate it with HandleBlank.
	u.base.HandleFunc("/", u.HandleBlank).Methods("GET")

	// Register this router with the baserouter package.
	baserouter.GetInstance().SetRouter("base", u.base)
}

// Start initializes the server, registers controllers, and starts listening.
func (u *muxServer) Start() {
	// Log server startup message with the configured address and port.
	log.Println("Server listening on ", config.GetInstance().Server.Address, ":", config.GetInstance().Server.Port)

	// Register controllers for handling API routes.
	basecontrollers.GetInstance().RegisterControllers()

	// Handle requests using the configured Mux router.
	http.Handle("/", u.base)

	// Start the HTTP server and listen on the configured address and port.
	err := http.ListenAndServe(config.GetInstance().Server.Address+":"+config.GetInstance().Server.Port, nil)
	if err != nil {
		log.Println("Error in running server:", err)
	}
}
