package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"websays/app/controllers"
	"websays/app/models"
	"websays/httpHandler/basecontrollers"

	"github.com/gorilla/mux"
)

func TestCreateArticle(t *testing.T) {

	controller, _ := basecontrollers.GetInstance().GetController("Article")
	articleController := controller.(*controllers.Article)
	// Create a sample JSON payload
	requestData := map[string]interface{}{
		"title": "firstArticle",
		"body":  "Article Body",
	}

	// Convert the JSON payload to a byte slice
	payload, err := json.Marshal(requestData)
	if err != nil {
		t.Fatal(err)
	}

	// Create a request with the JSON payload
	req, err := http.NewRequest("POST", "/createArticle", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call your API handler function
	articleController.HandleAddArticle(rr, req)
	// articleController. (rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, rr.Code)
	}

	// Parse and validate the response JSON (if applicable)
	var responseJSON map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&responseJSON)
	if err != nil {
		t.Fatal(err)
	}

	// Perform assertions on the response data (if applicable)
	expectedMessage := "Adding article success"
	if responseJSON["message"] != expectedMessage {
		t.Errorf("Expected message '%s'; got '%s'", expectedMessage, responseJSON["message"])
	}
}

func TestFindArticle(t *testing.T) {

	controller, _ := basecontrollers.GetInstance().GetController("Article")
	articleController := controller.(*controllers.Article)

	articleData := models.Article{
		Title: "First Article",
		Body:  "Articles body",
		ID:    articleController.GetNextID(),
	}
	articleController.GetFunctions()
	articleController.Add(articleController.GetDBName(), articleController.GetCollectionName(), articleData)

	// Create a new router
	r := mux.NewRouter()

	// Add the route with the same pattern used in your API
	r.HandleFunc("/api/readArticle/{id:[0-9]+}", articleController.HandleReadArticle)

	// Create a test server with the router
	server := httptest.NewServer(r)
	defer server.Close()

	// Generate a valid article ID (replace with your logic to generate an ID)
	articleID := strconv.FormatInt(int64(articleData.ID), 10)

	// Create the URL with the article ID
	url := server.URL + "/api/readArticle/" + articleID

	// Make a GET request to the test server with the generated URL
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Check the response status code (200 indicates success)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, resp.StatusCode)
	}

	var responseJSON map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&responseJSON)
	if err != nil {
		t.Fatal(err)
	}

	// Perform assertions on the response data (if applicable)
	expectedMessage := "Reading Article success"
	if responseJSON["message"] != expectedMessage {
		t.Errorf("Expected message '%s'; got '%s'", expectedMessage, responseJSON["message"])
	}

}

func TestDeleteArticle(t *testing.T) {

	controller, _ := basecontrollers.GetInstance().GetController("Article")
	articleController := controller.(*controllers.Article)

	articleData := models.Article{
		Title: "First Article",
		Body:  "Articles body",
		ID:    articleController.GetNextID(),
	}
	articleController.GetFunctions()
	articleController.Add(articleController.GetDBName(), articleController.GetCollectionName(), articleData)

	// Create a new router
	r := mux.NewRouter()

	// Add the route with the same pattern used in your API
	r.HandleFunc("/api/deleteArticle/{id:[0-9]+}", articleController.HandleDeleteArticle)

	// Create a test server with the router
	server := httptest.NewServer(r)
	defer server.Close()

	// Generate a valid article ID (replace with your logic to generate an ID)
	articleID := strconv.FormatInt(int64(articleData.ID), 10)

	// Create the URL with the article ID
	url := server.URL + "/api/deleteArticle/" + articleID
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Send the DELETE request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Check the response status code (200 indicates success)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, resp.StatusCode)
	}

	var responseJSON map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&responseJSON)
	if err != nil {
		t.Fatal(err)
	}

	// Perform assertions on the response data (if applicable)
	expectedMessage := "Deleting article success"
	if responseJSON["message"] != expectedMessage {
		t.Errorf("Expected message '%s'; got '%s'", expectedMessage, responseJSON["message"])
	}

}

func TestUpdateArticle(t *testing.T) {
	controller, _ := basecontrollers.GetInstance().GetController("Article")
	articleController := controller.(*controllers.Article)

	articleData := models.Article{
		Title: "First Article",
		Body:  "Articles body",
		ID:    articleController.GetNextID(),
	}
	articleController.GetFunctions()
	articleController.Add(articleController.GetDBName(), articleController.GetCollectionName(), articleData)

	// Create a sample JSON payload
	requestData := map[string]interface{}{
		"title": "firstArticle",
		"body":  "Article Body",
		"id":    articleData.ID,
	}

	// Convert the JSON payload to a byte slice
	payload, err := json.Marshal(requestData)
	if err != nil {
		t.Fatal(err)
	}

	// Create a request with the JSON payload
	req, err := http.NewRequest("POST", "/updateArticle", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call your API handler function
	articleController.HandleUpdateArticle(rr, req)
	// articleController. (rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, rr.Code)
	}

	// Parse and validate the response JSON (if applicable)
	var responseJSON map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&responseJSON)
	if err != nil {
		t.Fatal(err)
	}

	// Perform assertions on the response data (if applicable)
	expectedMessage := "Update article success"
	if responseJSON["message"] != expectedMessage {
		t.Errorf("Expected message '%s'; got '%s'", expectedMessage, responseJSON["message"])
	}
}
