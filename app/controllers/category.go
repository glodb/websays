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

// Category represents a controller for handling API endpoints related to catogeries.
// Because category in linked to file function, it will operate handling files
// It implements the controller interface and requires a Controller Factory,
// Base Functions Interface, and Validator Interface to operate effectively.
type Category struct {
	baseinterfaces.BaseControllerFactory
	basefunctions.BaseFucntionsInterface
	basevalidators.ValidatorInterface
}

// GetDBName returns the database name associated with the Category controller.
//
// This method retrieves the database name from the application configuration and
// returns it as a basetypes.DBName type.
//
// Parameters:
//   - None
//
// Returns:
//   - basetypes.DBName: The name of the database.
func (cat *Category) GetDBName() basetypes.DBName {
	return basetypes.DBName(config.GetInstance().Database.DBName)
}

// GetCollectionName returns the collection name associated with the Category controller.
//
// This method specifies the collection name for Category-related data as "categories."
//
// Parameters:
//   - None
//
// Returns:
//   - basetypes.CollectionName: The name of the collection.
func (cat *Category) GetCollectionName() basetypes.CollectionName {
	return "categories"
}

// DoIndexing performs any indexing operations required for the Category controller.
//
// This method is a no-op and returns nil, as no specific indexing operations are needed.
//
// Parameters:
//   - None
//
// Returns:
//   - error: Always returns nil.
func (cat *Category) DoIndexing() error {
	return nil
}

// SetBaseFunctions sets the BaseFunctionsInterface for the Category controller.
//
// This method allows the Category controller to set its BaseFunctionsInterface to
// facilitate interactions with the underlying data storage or controller.
//
// Parameters:
//   - inter: An instance of the BaseFunctionsInterface.
//
// Returns:
//   - None
func (cat *Category) SetBaseFunctions(inter basefunctions.BaseFucntionsInterface) {
	cat.BaseFucntionsInterface = inter
}

// HandleCreateCategory handles the creation of a new category based on the provided JSON data in the request body.
//
// This method expects a JSON object representing a category in the request body and validates it.
// If the JSON data is valid, it assigns a unique ID to the category, adds it to the underlying data storage
// using the Add method, and responds with a JSON-encoded success message.
//
// Parameters:
//   - w:   The http.ResponseWriter for sending the HTTP response.
//   - r:   The http.Request containing the incoming HTTP request.
//
// Behavior:
//   - Decodes the JSON data from the request body into a Category struct.
//   - Validates the category data using the Validate method.
//   - Assigns a unique ID to the category using the GetNextID method.
//   - Calls the Add method to add the category to the underlying data storage.
//   - Responds with a JSON-encoded success message upon successful creation.
//   - Responds with an error message if the JSON is malformed, validation fails, or the addition operation fails.
func (cat *Category) HandleCreateCategory(w http.ResponseWriter, r *http.Request) {
	category := models.Category{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&category)
	defer r.Body.Close()
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.MALFORMED_JSON, errors.New("illegal json format"), nil)
		return
	}

	err = cat.Validate(r.URL.Path, category)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	category.ID = cat.GetNextID()

	// Call the underlying file controller method
	_, err = cat.Add(cat.GetDBName(), cat.GetCollectionName(), category)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	responses.GetInstance().WriteJsonResponse(w, r, responses.ADD_CATEGORY_SUCCESS, nil, category)
}

// HandleReadCategory handles the retrieval of a category based on the provided ID in the request parameters.
//
// This method expects an ID as a route parameter in the URL, which is used to identify and retrieve the corresponding category.
// The category retrieval is performed by calling the FindOne method from the underlying data controller.
//
// Parameters:
//   - w:   The http.ResponseWriter for sending the HTTP response.
//   - r:   The http.Request containing the incoming HTTP request.
//
// Behavior:
//   - Extracts the category ID from the route parameters.
//   - Validates the category ID and converts it to an integer.
//   - Calls the FindOne method to retrieve the category data from the underlying data storage.
//   - Responds with a JSON-encoded success message containing the retrieved data upon successful retrieval.
//   - Responds with an error message if the validation or retrieval operation fails.
func (cat *Category) HandleReadCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	category := models.Category{}

	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	category.ID = int(idInt)

	// Call the underlying file controller find
	data, err := cat.FindOne(cat.GetDBName(), cat.GetCollectionName(), category)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	responses.GetInstance().WriteJsonResponse(w, r, responses.READ_CATEGORY_SUCCESS, nil, data)
}

// HandleUpdateCategory handles the update of a category based on the provided data in the request body.
//
// This method expects a JSON-encoded category object in the request body, which is used to update the corresponding category.
// The category update is performed by calling the UpdateOne method from the underlying data controller.
//
// Parameters:
//   - w:   The http.ResponseWriter for sending the HTTP response.
//   - r:   The http.Request containing the incoming HTTP request.
//
// Behavior:
//   - Parses the JSON-encoded category data from the request body.
//   - Validates the category data using the Validate method.
//   - Calls the UpdateOne method to update the category in the underlying data storage.
//   - Responds with a JSON-encoded success message containing the updated category data upon successful update.
//   - Responds with an error message if the validation or update operation fails.
func (cat *Category) HandleUpdateCategory(w http.ResponseWriter, r *http.Request) {
	category := models.Category{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&category)
	defer r.Body.Close()
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.MALFORMED_JSON, errors.New("illegal json format"), nil)
		return
	}

	// Call the validate method
	err = cat.Validate(r.URL.Path, category)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	// Call the underlying file controller update method
	err = cat.UpdateOne(cat.GetDBName(), cat.GetCollectionName(), "", category, false)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}
	responses.GetInstance().WriteJsonResponse(w, r, responses.UPDATE_CATEGORY_SUCCESS, nil, category)
}

// HandleDeleteCategory handles the deletion of a category based on the provided ID in the request.
//
// This method expects a category ID as a route parameter in the URL, which is used to identify and delete the corresponding category.
// The category deletion is performed by calling the DeleteOne method from the underlying data controller.
//
// Parameters:
//   - w:   The http.ResponseWriter for sending the HTTP response.
//   - r:   The http.Request containing the incoming HTTP request.
//
// Behavior:
//   - Extracts the category ID from the route parameters.
//   - Validates the category ID and converts it to an integer.
//   - Calls the DeleteOne method to delete the category in the underlying data storage.
//   - Responds with a JSON-encoded success message upon successful deletion.
//   - Responds with an error message if the validation or deletion operation fails.
func (cat *Category) HandleDeleteCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	category := models.Category{}

	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	category.ID = int(idInt)

	// Call the underlying file delete controller
	err = cat.DeleteOne(cat.GetDBName(), cat.GetCollectionName(), category)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}
	responses.GetInstance().WriteJsonResponse(w, r, responses.DELETE_CATEGORY_SUCCESS, nil, category)
}

// RegisterApis registers the API endpoints associated with category operations.
//
// This method configures the routing for category-related API endpoints using the provided base router.
// It maps HTTP methods to the corresponding category handler functions:
//   - POST   -> /api/createCategory: HandleCreateCategory
//   - GET    -> /api/readCategory/{id}: HandleReadCategory
//   - PUT    -> /api/updateCategory: HandleUpdateCategory
//   - DELETE -> /api/deleteCategory/{id}: HandleDeleteCategory
//
// Behavior:
//   - Configures HTTP routes for category-related operations.
//   - Associates each route with its respective handler function.
func (cat *Category) RegisterApis() {
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/createCategory", cat.HandleCreateCategory).Methods("POST")
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/readCategory/{id}", cat.HandleReadCategory).Methods("GET")
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/updateCategory", cat.HandleUpdateCategory).Methods("PUT")
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/deleteCategory/{id}", cat.HandleDeleteCategory).Methods("DELETE")
}
