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

type muxServer struct {
	base *mux.Router
}

var (
	instance *muxServer
	once     sync.Once
)

//Singleton. Returns a single object
func GetInstance() *muxServer {
	// var instance
	once.Do(func() {
		instance = &muxServer{}
		instance.setup()
	})
	return instance
}

func (u *muxServer) HandleBlank(w http.ResponseWriter, r *http.Request) {
	responses.GetInstance().WriteJsonResponse(w, r, responses.WEBSAYS_TEST, nil, nil)
}

func (u *muxServer) setup() {
	corsMiddleware := middlewares.CORSMiddleware{}
	u.base = &mux.Router{}

	u.base.Use(corsMiddleware.GetHandlerFunc)
	u.base.HandleFunc("/", u.HandleBlank).Methods("GET")

	baserouter.GetInstance().SetRouter("base", u.base)
}

func (u *muxServer) Start() {
	log.Println("Server listening on ", config.GetInstance().Server.Address, ":", config.GetInstance().Server.Port)
	basecontrollers.GetInstance().RegisterControllers()

	http.Handle("/", u.base)
	err := http.ListenAndServe(config.GetInstance().Server.Address+":"+config.GetInstance().Server.Port, nil)
	if err != nil {
		log.Println("Error in running server:", err)
	}
}
