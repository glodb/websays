package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"websays/app/models"
	"websays/config"
	"websays/database/basefunctions"
	"websays/database/basetypes"
	"websays/httpHandler/basecontrollers/baseinterfaces"
	"websays/httpHandler/baserouter"
	"websays/httpHandler/basevalidators"
	"websays/httpHandler/responses"

	"github.com/gorilla/mux"
)

// Article represents a controller for handling API endpoints related to articles.
// It implements the controller interface and requires a Controller Factory,
// Base Functions Interface, and Validator Interface to operate effectively.
// Articles are linked to memory functions they work with memory
type Article struct {
	baseinterfaces.BaseControllerFactory
	basefunctions.BaseFucntionsInterface
	basevalidators.ValidatorInterface
}

// GetDBName returns the database name associated with the Article controller.
//
// Input:
//   None.
//
// Output:
//   - basetypes.DBName: The database name as a basetypes.DBName type.
func (art *Article) GetDBName() basetypes.DBName {
	return basetypes.DBName(config.GetInstance().Database.DBName)
}

// GetCollectionName returns the collection name associated with the Article controller.
//
// Input:
//   None.
//
// Output:
//   - basetypes.CollectionName: The collection name as a basetypes.CollectionName type.
func (art *Article) GetCollectionName() basetypes.CollectionName {
	return "articles"
}

// DoIndexing performs indexing-related operations associated with the Article controller.
//
// This method is currently implemented to return nil, indicating that no specific
// indexing operations are performed by the Article controller. Implementations of
// this method can include custom logic for indexing if needed in the future for memory controllers.
//
// Returns:
//   - error: An error is returned if there are any issues with indexing operations.
//             In the current implementation, it always returns nil.
func (art *Article) DoIndexing() error {
	return nil
}

// SetBaseFunctions sets the implementation of the BaseFunctionsInterface for the Article controller.
//
// This method allows assigning a specific implementation of the BaseFunctionsInterface
// to the Article controller. The BaseFunctionsInterface provides functionality related
// to basic CRUD (Create, Read, Update, Delete) operations on data.
//
// Parameters:
//   - inter: An instance of basefunctions.BaseFucntionsInterface to be set as the
//            implementation for the Article controller.
func (art *Article) SetBaseFunctions(inter basefunctions.BaseFucntionsInterface) {
	art.BaseFucntionsInterface = inter
}

// HandleAddArticle handles the creation of a new article based on the JSON data provided in the request body.
//
// This method parses the incoming JSON data into an article structure and validates it using the article validator.
// If the JSON data is malformed or validation fails, it responds with an appropriate error message.
// If validation succeeds, it assigns a new unique ID to the article and adds it to the underlying memory controller using the Add method.
// Finally, it responds with a JSON-encoded success message along with the created article.
//
// Parameters:
//   - w:   The http.ResponseWriter for sending the HTTP response.
//   - r:   The http.Request containing the incoming HTTP request.
//
// Behavior:
//   - Decodes the JSON data from the request body into an article structure.
//   - Validates the article data using the article validator.
//   - Generates a new unique ID for the article.
//   - Adds the article to the underlying memory controller using the Add method.
//   - Responds with a JSON-encoded success message and the created article upon successful addition.
//   - Responds with an error message if the JSON data is malformed or validation fails.
func (art *Article) HandleAddArticle(w http.ResponseWriter, r *http.Request) {
	article := models.Article{}

	// Decode the JSON data from the request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&article)
	defer r.Body.Close()
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.MALFORMED_JSON, errors.New("illegal json format"), nil)
		return
	}

	// Validate the article data using the article validator
	err = art.Validate(r.URL.Path, article)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	// Generate a new unique ID for the article
	article.ID = art.GetNextID()

	// Add the article to the underlying memory controller
	err = art.Add(art.GetDBName(), art.GetCollectionName(), article)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	// Respond with a JSON-encoded success message and the created article
	responses.GetInstance().WriteJsonResponse(w, r, responses.ADD_ARTICLE_SUCCESS, nil, article)
}

// HandleReadArticle handles the retrieval of an article based on the provided ID in the request.
//
// This method expects an article ID as a route parameter in the URL, which is used to identify and
// retrieve the corresponding article. The article retrieval is performed by calling the FindOne method
// from the underlying memory controller.
//
// Parameters:
//   - w:   The http.ResponseWriter for sending the HTTP response.
//   - r:   The http.Request containing the incoming HTTP request.
//
// Behavior:
//   - Extracts the article ID from the route parameters.
//   - Validates the article ID and converts it to an integer.
//   - Calls the FindOne method to retrieve the article in the underlying memory controller.
//   - Responds with a JSON-encoded article data upon successful retrieval.
//   - Responds with an error message if the validation or retrieval operation fails.
func (art *Article) HandleReadArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	article := models.Article{}

	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	article.ID = int(idInt)

	// Calling the FindOne method for the memory controller
	data, err := art.FindOne(art.GetDBName(), art.GetCollectionName(), article)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	// Respond with a JSON-encoded article data
	responses.GetInstance().WriteJsonResponse(w, r, responses.READ_ARTICLE_SUCCESS, nil, data)
}

// HandleDeleteArticle handles the deletion of an article based on the provided ID in the request.
//
// This method expects an article ID as a route parameter in the URL, which is used to identify and
// delete the corresponding article. The article deletion is performed by calling the DeleteOne method
// from the underlying memory controller.
//
// Parameters:
//   - w:   The http.ResponseWriter for sending the HTTP response.
//   - r:   The http.Request containing the incoming HTTP request.
//
// Behavior:
//   - Extracts the article ID from the route parameters.
//   - Validates the article ID and converts it to an integer.
//   - Calls the DeleteOne method to delete the article in the underlying memory controller.
//   - Responds with a JSON-encoded success message upon successful deletion.
//   - Responds with an error message if the validation or deletion operation fails.
func (art *Article) HandleDeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	article := models.Article{}

	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	article.ID = int(idInt)

	// Calling the DeleteOne method for the memory controller
	err = art.DeleteOne(art.GetDBName(), art.GetCollectionName(), article)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	// Respond with a JSON-encoded success message
	responses.GetInstance().WriteJsonResponse(w, r, responses.DELETE_ARTICLE_SUCCESS, nil, nil)
}

// HandleUpdateArticle handles the update of an article based on the provided JSON data in the request.
//
// This method expects a JSON-encoded article in the request body, which is decoded and used to update
// an existing article's details. The article update is performed by calling the UpdateOne method
// from the underlying memory controller.
//
// Parameters:
//   - w:   The http.ResponseWriter for sending the HTTP response.
//   - r:   The http.Request containing the incoming HTTP request.
//
// Behavior:
//   - Decodes the JSON data from the request body into an article structure.
//   - Validates the article data using the Validate method.
//   - Calls the UpdateOne method to update the article in the underlying memory controller.
//   - Responds with a JSON-encoded success message and the updated article upon success.
//   - Responds with an error message if decoding, validation, or the update operation fails.
func (art *Article) HandleUpdateArticle(w http.ResponseWriter, r *http.Request) {
	article := models.Article{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&article)
	defer r.Body.Close()
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.MALFORMED_JSON, errors.New("illegal JSON format"), nil)
		return
	}

	err = art.Validate(r.URL.Path, article)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	// Calling the UpdateOne method for the underlying memory controller
	err = art.UpdateOne(art.GetDBName(), art.GetCollectionName(), "", article, false)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	// Respond with a JSON-encoded success message and the updated article
	responses.GetInstance().WriteJsonResponse(w, r, responses.UPDATE_ARTICLE_SUCCESS, nil, article)
}

// RegisterApis registers the API endpoints associated with the Article controller.
//
// This method configures the routes and HTTP methods for various Article-related actions:
//   - /api/createArticle:   Handles the creation of a new article (HTTP POST).
//   - /api/readArticle/{id}: Handles the retrieval of an article by ID (HTTP GET).
//   - /api/deleteArticle/{id}: Handles the deletion of an article by ID (HTTP DELETE).
//   - /api/updateArticle:   Handles the update of an existing article (HTTP PUT).
//
// Parameters:
//   - None
//
// Behavior:
//   - Configures the routes and handlers for the specified API endpoints.
func (art *Article) RegisterApis() {
	// Register API endpoints with their respective handlers and HTTP methods
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/createArticle", art.HandleAddArticle).Methods("POST")
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/readArticle/{id}", art.HandleReadArticle).Methods("GET")
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/deleteArticle/{id}", art.HandleDeleteArticle).Methods("DELETE")
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/updateArticle", art.HandleUpdateArticle).Methods("PUT")
}
