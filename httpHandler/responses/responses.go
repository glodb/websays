package responses

import (
	"encoding/json"
	"net/http"
	"sync"
)

const (
	WEBSAYS_TEST            = 1000
	API_NOT_AVAILABLE       = 1001
	OPTIONS_NOT_ALLOWED     = 1002
	MALFORMED_JSON          = 1003
	VALIDATION_FAILED       = 1004
	ADDING_DB_FAILED        = 1005
	ADD_ARTICLE_SUCCESS     = 1006
	READ_ARTICLE_SUCCESS    = 1007
	DELETE_ARTICLE_SUCCESS  = 1008
	UPDATE_ARTICLE_SUCCESS  = 1009
	ADD_CATEGORY_SUCCESS    = 1010
	READ_CATEGORY_SUCCESS   = 1011
	UPDATE_CATEGORY_SUCCESS = 1012
	DELETE_CATEGORY_SUCCESS = 1013
	ADD_PRODUCT_SUCCESS     = 1014
	READ_PRODUCT_SUCCESS    = 1015
	UPDATE_PRODUCT_SUCCESS  = 1016
	DELETE_PRODUCT_SUCCESS  = 1017
	NO_PRDUCT_FOUND         = 1018
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
	u.responses[ADD_CATEGORY_SUCCESS] = "Adding category success"
	u.responses[READ_CATEGORY_SUCCESS] = "Reading category success"
	u.responses[UPDATE_CATEGORY_SUCCESS] = "Updating category success"
	u.responses[DELETE_CATEGORY_SUCCESS] = "Deleting category success"
	u.responses[ADD_PRODUCT_SUCCESS] = "Adding product success"
	u.responses[READ_PRODUCT_SUCCESS] = "Reading product success"
	u.responses[UPDATE_PRODUCT_SUCCESS] = "Updating product success"
	u.responses[DELETE_PRODUCT_SUCCESS] = "Deleting product success"
	u.responses[NO_PRDUCT_FOUND] = "No product found"
}

// GetResponse returns the message for the particular response code
func (u *Responses) getResponse(code int) map[string]interface{} {
	message := make(map[string]interface{})
	message["code"] = code
	message["message"] = u.responses[code]

	return message
}

// WriteJsonResponse writes a JSON response to the HTTP response writer based on the provided parameters.
// It sets the appropriate HTTP status code, content type, and includes data or error information in the response.
//
// Parameters:
//   - w: The http.ResponseWriter to which the JSON response will be written.
//   - r: The *http.Request containing information about the incoming request.
//   - code: The HTTP status code to be set in the response.
//   - err: An optional error to be included in the response if applicable. If nil, no error is included.
//   - data: An optional data payload to be included in the response if applicable. If nil, no data is included.
//
// Note:
//   - The 'err' parameter is used to indicate if there is an error associated with the response, and it affects the
//     HTTP status code. If 'err' is not nil, the status code is set to StatusNotAcceptable (406); otherwise, it's
//     set to StatusOK (200).
//   - The response format is JSON with appropriate headers.
//   - If encoding the JSON response encounters an error, it responds with an internal server error (HTTP status 500).
//
// Example Usage:
//   Responses{}.WriteJsonResponse(w, r, http.StatusOK, nil, responseData)
//   Responses{}.WriteJsonResponse(w, r, http.StatusNotAcceptable, validationError, nil)
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
