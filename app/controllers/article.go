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

type Article struct {
	baseinterfaces.BaseControllerFactory
	basefunctions.BaseFucntionsInterface
	basevalidators.ValidatorInterface
	datastore map[int]interface{}
	id        int
}

func (art *Article) GetDBName() basetypes.DBName {
	return basetypes.DBName(config.GetInstance().Database.DBName)
}

func (art *Article) GetCollectionName() basetypes.CollectionName {
	return "articles"
}

func (art *Article) DoIndexing() error {
	return nil
}

func (art *Article) SetBaseFunctions(inter basefunctions.BaseFucntionsInterface) {
	art.BaseFucntionsInterface = inter
}

func (art *Article) HandleAddArticle(w http.ResponseWriter, r *http.Request) {
	article := models.Article{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&article)
	defer r.Body.Close()
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.MALFORMED_JSON, errors.New("illegal json format"), nil)
		return
	}

	err = art.Validate(r.URL.Path, article)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	article.ID = art.GetNextID()

	err = art.Add(art.GetDBName(), art.GetCollectionName(), article)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	responses.GetInstance().WriteJsonResponse(w, r, responses.ADD_ARTICLE_SUCCESS, nil, article)
}

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

	data, err := art.FindOne(art.GetDBName(), art.GetCollectionName(), article)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	responses.GetInstance().WriteJsonResponse(w, r, responses.READ_ARTICLE_SUCCESS, nil, data)
}

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

	err = art.DeleteOne(art.GetDBName(), art.GetCollectionName(), article)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	responses.GetInstance().WriteJsonResponse(w, r, responses.DELETE_ARTICLE_SUCCESS, nil, nil)
}

func (art *Article) HandleUpdateArticle(w http.ResponseWriter, r *http.Request) {
	article := models.Article{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&article)
	defer r.Body.Close()
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.MALFORMED_JSON, errors.New("illegal json format"), nil)
		return
	}

	err = art.Validate(r.URL.Path, article)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}

	err = art.UpdateOne(art.GetDBName(), art.GetCollectionName(), "", article, false)
	if err != nil {
		responses.GetInstance().WriteJsonResponse(w, r, responses.VALIDATION_FAILED, err, nil)
		return
	}
	responses.GetInstance().WriteJsonResponse(w, r, responses.UPDATE_ARTICLE_SUCCESS, nil, article)
}

func (art *Article) RegisterApis() {
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/createArticle", art.HandleAddArticle).Methods("POST")
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/readArticle/{id}", art.HandleReadArticle).Methods("GET")
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/deleteArticle/{id}", art.HandleDeleteArticle).Methods("DELETE")
	baserouter.GetInstance().GetBaseRouter().HandleFunc("/api/updateArticle", art.HandleUpdateArticle).Methods("PUT")
}
