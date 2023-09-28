package responses

import (
	"encoding/json"
	"net/http"
	"sync"
)

const (
	WEBSAYS_TEST           = 1000
	API_NOT_AVAILABLE      = 1001
	OPTIONS_NOT_ALLOWED    = 1002
	MALFORMED_JSON         = 1003
	VALIDATION_FAILED      = 1004
	ADDING_DB_FAILED       = 1005
	ADD_ARTICLE_SUCCESS    = 1006
	READ_ARTICLE_SUCCESS   = 1007
	DELETE_ARTICLE_SUCCESS = 1008
	UPDATE_ARTICLE_SUCCESS = 1009
)

type Responses struct {
	responses map[int]string
}

var (
	instance *Responses
	once     sync.Once
)

// Singleton. Returns a single object of Factory
func GetInstance() *Responses {
	// var instance
	once.Do(func() {
		instance = &Responses{}
		instance.InitResponses()
	})
	return instance
}

// InitResponses function just initialise the response headers to be sent
func (u *Responses) InitResponses() {
	u.responses = make(map[int]string)
	u.responses[WEBSAYS_TEST] = "Websays API's"
	u.responses[API_NOT_AVAILABLE] = "The Api is not available on current server"
	u.responses[OPTIONS_NOT_ALLOWED] = "Options are not allowed"
	u.responses[MALFORMED_JSON] = "Json is not correctly formed"
	u.responses[VALIDATION_FAILED] = "Failure in validation"
	u.responses[ADD_ARTICLE_SUCCESS] = "Adding article success"
	u.responses[ADDING_DB_FAILED] = "Adding DB Failed"
	u.responses[READ_ARTICLE_SUCCESS] = "Reading Article success"
	u.responses[DELETE_ARTICLE_SUCCESS] = "Deleting article success"
	u.responses[UPDATE_ARTICLE_SUCCESS] = "Update article success"
}

// GetResponse returns the message for the particular response code
func (u *Responses) getResponse(code int) map[string]interface{} {
	message := make(map[string]interface{})
	message["code"] = code
	message["message"] = u.responses[code]

	return message
}

func (u *Responses) WriteJsonResponse(w http.ResponseWriter, r *http.Request, code int, err error, data interface{}) {
	// urlPath := r.URL
	returnMap := u.getResponse(code)

	status := http.StatusOK
	if err != nil {
		status = http.StatusNotAcceptable
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")

	w.WriteHeader(status)

	if err != nil {
		returnMap["error"] = err.Error()
	}
	if data != nil {
		returnMap["data"] = data
	}
	err = json.NewEncoder(w).Encode(returnMap)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
