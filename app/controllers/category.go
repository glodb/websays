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

type Category struct {
	baseinterfaces.BaseControllerFactory
	basefunctions.BaseFucntionsInterface
	basevalidators.ValidatorInterface
}

func (cat *Category) GetDBName() basetypes.DBName {
	return basetypes.DBName(config.GetInstance().Database.DBName)
}

func (cat *Category) GetCollectionName() basetypes.CollectionName {
	return "categories"
}

func (cat *Category) DoIndexing() error {
	return nil
}

func (cat *Category) SetBaseFunctions(inter basefunctions.BaseFucntionsInterface) {
	cat.BaseFucntionsInterface = inter
}

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

	err = cat.Add(cat.GetDBName(), cat.GetCollectionName(), category)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	responses.GetInstance().WriteJsonResponse(w, r, responses.ADD_CATEGORY_SUCCESS, nil, category)
}

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

	data, err := cat.FindOne(cat.GetDBName(), cat.GetCollectionName(), category)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	responses.GetInstance().WriteJsonResponse(w, r, responses.READ_CATEGORY_SUCCESS, nil, data)
}

func (cat *Category) HandleUpdateCategory(w http.ResponseWriter, r *http.Request) {
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

	err = cat.UpdateOne(cat.GetDBName(), cat.GetCollectionName(), "", category, false)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}
	responses.GetInstance().WriteJsonResponse(w, r, responses.UPDATE_CATEGORY_SUCCESS, nil, category)
}

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

	err = cat.DeleteOne(cat.GetDBName(), cat.GetCollectionName(), category)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}
	responses.GetInstance().WriteJsonResponse(w, r, responses.DELETE_CATEGORY_SUCCESS, nil, category)
}
func (cat *Category) RegisterApis() {
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/createCategory", cat.HandleCreateCategory).Methods("POST")
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/readCategory/{id}", cat.HandleReadCategory).Methods("GET")
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/updateCategory", cat.HandleUpdateCategory).Methods("PUT")
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/deleteCategory/{id}", cat.HandleDeleteCategory).Methods("DELETE")
}
