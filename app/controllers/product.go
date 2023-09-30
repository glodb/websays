package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
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

// Product represents a controller for handling API endpoints related to products.
// Because product in linked to mysql functions, it will operate handling db.
// It implements the controller interface and requires a Controller Factory,
// Base Functions Interface, and Validator Interface to operate effectively.
type Product struct {
	baseinterfaces.BaseControllerFactory
	basefunctions.BaseFucntionsInterface
	basevalidators.ValidatorInterface
	isIndexed bool
}

// GetDBName returns the database name for product-related operations.
//
// This method retrieves the database name from the configuration instance and returns it as a DBName.
//
// Returns:
//   - DBName: The name of the database used for product-related operations.
func (pro *Product) GetDBName() basetypes.DBName {
	return basetypes.DBName(config.GetInstance().Database.DBName)
}

// GetCollectionName returns the collection name for product-related operations.
//
// This method returns the name of the collection used for product-related operations.
//
// Returns:
//   - CollectionName: The name of the collection used for product-related operations.
func (pro *Product) GetCollectionName() basetypes.CollectionName {
	return "products"
}

// DoIndexing performs indexing for product-related data.
//
// This method ensures the indexing of product data in the specified database and collection.
// It checks if the indexing process was successful and sets the "isIndexed" flag accordingly.
//
// Returns:
//   - error: An error if the indexing process encounters any issues; otherwise, nil.
func (pro *Product) DoIndexing() error {
	err := pro.EnsureIndex(pro.GetDBName(), pro.GetCollectionName(), models.Product{})
	if err == nil {
		pro.isIndexed = true
	}
	return nil
}

// SetBaseFunctions sets the base functions interface for product-related operations.
//
// This method assigns the provided base functions interface to the product controller.
//
// Parameters:
//   - inter: The base functions interface to be assigned.
func (pro *Product) SetBaseFunctions(inter basefunctions.BaseFucntionsInterface) {
	pro.BaseFucntionsInterface = inter
}

// HandleCreateProduct handles the creation of a new product based on the JSON data provided in the request body.
//
// This method performs the following steps:
//   - Checks if the product data is indexed; if not, it triggers the indexing process.
//   - Decodes the JSON data from the request body into a product struct.
//   - Validates the product data using the product validator.
//   - Adds the product to the database using the specified MySQL controller.
//   - Responds with a JSON-encoded success message upon successful product creation.
//   - Responds with an error message if the validation or creation operation fails.
//
// Parameters:
//   - w:   The http.ResponseWriter for sending the HTTP response.
//   - r:   The http.Request containing the incoming HTTP request.
func (pro *Product) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	if !pro.isIndexed {
		pro.DoIndexing()
	}

	product := models.Product{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&product)
	defer r.Body.Close()
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.MALFORMED_JSON, errors.New("illegal json format"), nil)
		return
	}

	// Calling the validate method of the product validator
	err = pro.Validate(r.URL.Path, product)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	// Calling the Add method for the MySQL controller
	product.ID, err = pro.Add(pro.GetDBName(), pro.GetCollectionName(), product)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	// Respond with a JSON-encoded success message
	responses.GetInstance().WriteJsonResponse(w, r, responses.ADD_PRODUCT_SUCCESS, nil, product)
}

// HandleReadProduct retrieves product information based on the provided product ID in the URL route parameter.
//
// This method performs the following steps:
//   - Extracts the product ID from the route parameters and validates it.
//   - Constructs a query to find a product with the specified ID.
//   - Calls the FindOne method for the MySQL controller to retrieve product data.
//   - Parses and extracts product details from the SQL result rows.
//   - Responds with a JSON-encoded success message containing the product information.
//   - Responds with an error message if the validation, query execution, or product not found.
//
// Parameters:
//   - w:   The http.ResponseWriter for sending the HTTP response.
//   - r:   The http.Request containing the incoming HTTP request.
func (pro *Product) HandleReadProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	query := make(map[string]interface{})
	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	query["id"] = int(idInt)

	// Calling the FindOne method for the MySQL controller
	data, err := pro.FindOne(pro.GetDBName(), pro.GetCollectionName(), query)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	rows := data.(*sql.Rows)

	var product models.Product
	rowsCount := 0

	// Reading and parsing rows
	for rows.Next() {
		err = rows.Scan(&product.ID, &product.Name)
		rowsCount++
	}
	log.Println(err)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	log.Println(rowsCount)
	if rowsCount == 0 {
		responses.GetInstance().WriteJsonResponse(w, r, responses.NO_PRDUCT_FOUND, errors.New("No Product found"), nil)
		return
	}

	// Respond with a JSON-encoded success message
	responses.GetInstance().WriteJsonResponse(w, r, responses.READ_PRODUCT_SUCCESS, nil, product)
}

// HandleUpdateProduct updates a product's information based on the provided JSON data in the HTTP request body.
//
// This method performs the following steps:
//   - Decodes the JSON data from the request body into a product struct.
//   - Calls the Validate method to validate the product data.
//   - Constructs conditions and data maps for updating the product.
//   - Calls the UpdateOne method for the underlying database controller to update the product.
//   - Responds with a JSON-encoded success message containing the updated product information.
//   - Responds with an error message if the JSON decoding, validation, or database update fails.
//
// Parameters:
//   - w:   The http.ResponseWriter for sending the HTTP response.
//   - r:   The http.Request containing the incoming HTTP request.
func (pro *Product) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	product := models.Product{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&product)
	defer r.Body.Close()
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.MALFORMED_JSON, errors.New("illegal json format"), nil)
		return
	}

	// Calling the Validate method to validate the product data
	err = pro.Validate(r.URL.Path, product)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	conditions := make(map[string]interface{})
	conditions["id"] = product.ID

	data := make(map[string]interface{})
	data["name"] = product.Name

	// Calling the UpdateOne method for the underlying database controller
	err = pro.UpdateOne(pro.GetDBName(), pro.GetCollectionName(), conditions, data, false)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	// Respond with a JSON-encoded success message
	responses.GetInstance().WriteJsonResponse(w, r, responses.UPDATE_PRODUCT_SUCCESS, nil, product)
}

// HandleDeleteProduct handles the deletion of a product based on the provided product ID in the route parameters.
//
// This method performs the following steps:
//   - Extracts the product ID from the route parameters.
//   - Validates the product ID and converts it to an integer.
//   - Constructs a conditions map for specifying the product to delete.
//   - Calls the DeleteOne method for the underlying database controller to delete the product.
//   - Responds with a JSON-encoded success message upon successful deletion.
//   - Responds with an error message if the validation or deletion operation fails.
//
// Parameters:
//   - w:   The http.ResponseWriter for sending the HTTP response.
//   - r:   The http.Request containing the incoming HTTP request.
func (pro *Product) HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	conditions := make(map[string]interface{})
	conditions["id"] = idInt

	// Calling the DeleteOne method for the underlying database controller
	err = pro.DeleteOne(pro.GetDBName(), pro.GetCollectionName(), conditions)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	// Respond with a JSON-encoded success message
	responses.GetInstance().WriteJsonResponse(w, r, responses.DELETE_PRODUCT_SUCCESS, nil, nil)
}

// RegisterApis registers the API endpoints for product-related operations.
//
// This method associates the HTTP handlers for creating, reading, updating, and deleting products
// with their respective API routes. It sets up the following API routes:
//   - POST /api/createProduct: Create a new product.
//   - GET /api/readProduct/{id}: Read the details of a product by its ID.
//   - DELETE /api/deleteProduct/{id}: Delete a product by its ID.
//   - PUT /api/updateProduct: Update the details of a product.
//
// Each API route is associated with a corresponding HTTP handler method in the Product controller.
//
// Usage:
//   Call this method during the application setup to define the product-related API endpoints.
func (pro *Product) RegisterApis() {
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/createProduct", pro.HandleCreateProduct).Methods("POST")
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/readProduct/{id}", pro.HandleReadProduct).Methods("GET")
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/deleteProduct/{id}", pro.HandleDeleteProduct).Methods("DELETE")
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/updateProduct", pro.HandleUpdateProduct).Methods("PUT")
}
