package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"websays/app/controllers"
	"websays/config"
	"websays/httpHandler/basecontrollers"

	"github.com/gorilla/mux"
)

func TestCreateCategory(t *testing.T) {

	config.GetInstance().FilePath = "/home/aafaq/Work/websays/testFiles"
	config.GetInstance().RunningFileName = ".runningNumber"
	controller, _ := basecontrollers.GetInstance().GetController("Category")
	categoryController := controller.(*controllers.Category)
	// Create a sample JSON payload
	requestData := map[string]interface{}{
		"name": "firstCategory",
		"id":   categoryController.GetNextID(),
	}

	// Convert the JSON payload to a byte slice
	payload, err := json.Marshal(requestData)
	if err != nil {
		t.Fatal(err)
	}

	// Create a request with the JSON payload
	req, err := http.NewRequest("POST", "/createCategory", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call your API handler function
	categoryController.HandleCreateCategory(rr, req)
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
	code := 1010
	if int(responseJSON["code"].(float64)) != code {
		t.Errorf("Expected message '%d'; got '%v'", code, responseJSON)
	}
}

func TestFindCategory(t *testing.T) {

	config.GetInstance().FilePath = "/home/aafaq/Work/websays/testFiles"
	config.GetInstance().RunningFileName = ".runningNumber"

	controller, _ := basecontrollers.GetInstance().GetController("Category")
	categoryController := controller.(*controllers.Category)

	// Create a new router
	r := mux.NewRouter()

	// Add the route with the same pattern used in your API
	r.HandleFunc("/api/readArticle/{id:[0-9]+}", categoryController.HandleReadCategory)

	// Create a test server with the router
	server := httptest.NewServer(r)
	defer server.Close()

	// Generate a valid article ID (replace with your logic to generate an ID)
	articleID := "2"

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
	code := 1011
	if int(responseJSON["code"].(float64)) != code {
		t.Errorf("Expected message '%d'; got '%v'", code, responseJSON)
	}

}

func TestDeleteCategory(t *testing.T) {

	config.GetInstance().FilePath = "/home/aafaq/Work/websays/testFiles"
	config.GetInstance().RunningFileName = ".runningNumber"

	controller, _ := basecontrollers.GetInstance().GetController("Category")
	categoryController := controller.(*controllers.Category)

	// Create a new router
	r := mux.NewRouter()

	// Add the route with the same pattern used in your API
	r.HandleFunc("/api/deleteArticle/{id:[0-9]+}", categoryController.HandleDeleteCategory)

	// Create a test server with the router
	server := httptest.NewServer(r)
	defer server.Close()

	// Generate a valid article ID (replace with your logic to generate an ID)
	articleID := "4"

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

	code := 1013
	if int(responseJSON["code"].(float64)) != code {
		t.Errorf("Expected message '%d'; got '%v'", code, responseJSON)
	}

}

func TestUpdateCategory(t *testing.T) {

	config.GetInstance().FilePath = "/home/aafaq/Work/websays/testFiles"
	config.GetInstance().RunningFileName = ".runningNumber"

	controller, _ := basecontrollers.GetInstance().GetController("Category")
	categoryController := controller.(*controllers.Category)

	// Create a sample JSON payload
	requestData := map[string]interface{}{
		"name": "updated",
		"id":   8,
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
	categoryController.HandleUpdateCategory(rr, req)
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

	code := 1012
	if int(responseJSON["code"].(float64)) != code {
		t.Errorf("Expected message '%d'; got '%v'", code, responseJSON)
	}
}
