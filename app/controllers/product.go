package controllers

import (
	"database/sql"
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

type Product struct {
	baseinterfaces.BaseControllerFactory
	basefunctions.BaseFucntionsInterface
	basevalidators.ValidatorInterface
	isIndexed bool
}

func (pro *Product) GetDBName() basetypes.DBName {
	return basetypes.DBName(config.GetInstance().Database.DBName)
}

func (pro *Product) GetCollectionName() basetypes.CollectionName {
	return "products"
}

func (pro *Product) DoIndexing() error {
	err := pro.EnsureIndex(pro.GetDBName(), pro.GetCollectionName(), models.Product{})
	if err == nil {
		pro.isIndexed = true
	}
	return nil
}

func (pro *Product) SetBaseFunctions(inter basefunctions.BaseFucntionsInterface) {
	pro.BaseFucntionsInterface = inter
}

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

	err = pro.Validate(r.URL.Path, product)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	err = pro.Add(pro.GetDBName(), pro.GetCollectionName(), product)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	responses.GetInstance().WriteJsonResponse(w, r, responses.ADD_PRODUCT_SUCCESS, nil, product)
}

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

	data, err := pro.FindOne(pro.GetDBName(), pro.GetCollectionName(), query)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	rows := data.(*sql.Rows)

	var product models.Product
	rowsCount := 0

	for rows.Next() {

		err = rows.Scan(&product.ID, &product.Name)
	}
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	if rowsCount == 0 {
		responses.GetInstance().WriteJsonResponse(w, r, responses.NO_PRDUCT_FOUND, errors.New("No Product found"), nil)
		return
	}

	responses.GetInstance().WriteJsonResponse(w, r, responses.READ_PRODUCT_SUCCESS, nil, product)
}

func (pro *Product) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	product := models.Product{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&product)
	defer r.Body.Close()
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.MALFORMED_JSON, errors.New("illegal json format"), nil)
		return
	}

	err = pro.Validate(r.URL.Path, product)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	conditions := make(map[string]interface{})
	conditions["id"] = product.ID

	data := make(map[string]interface{})
	data["name"] = product.Name
	err = pro.UpdateOne(pro.GetDBName(), pro.GetCollectionName(), conditions, data, false)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}
	responses.GetInstance().WriteJsonResponse(w, r, responses.UPDATE_PRODUCT_SUCCESS, nil, product)
}

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

	err = pro.DeleteOne(pro.GetDBName(), pro.GetCollectionName(), conditions)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	responses.GetInstance().WriteJsonResponse(w, r, responses.DELETE_PRODUCT_SUCCESS, nil, nil)
}
func (pro *Product) RegisterApis() {
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/createProduct", pro.HandleCreateProduct).Methods("POST")
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/readProduct/{id}", pro.HandleReadProduct).Methods("GET")
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/deleteProduct/{id}", pro.HandleDeleteProduct).Methods("DELETE")
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/updateProduct", pro.HandleUpdateProduct).Methods("PUT")
}
