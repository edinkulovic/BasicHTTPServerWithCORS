package routes

import (
	"fmt"
	"net/http"
	"encoding/json"
)

// Server implementation of all routes
type Server struct{}

// Routes definitions
type Routes struct {}

// Mux route definitions
var Mux map[string]http.HandlerFunc

func init() {
	Mux = make(map[string]http.HandlerFunc)

	// Define your routes here
	Mux["/"] =Routes{}.HealtCheck
	Mux["/test_post"] = Routes{}.TestPost
	Mux["/test_get"] = Routes{}.TestGet
}

// HealtCheck method is used to check if server is live.
func (r Routes) HealtCheck(writer http.ResponseWriter, request *http.Request) {
	return
}

// Test is temporary structure 
type Test struct {
	Text    string  `json:"Text"`

}

// TestPost will just log something from the form
func (r Routes) TestPost(w http.ResponseWriter, request *http.Request) {
	// Important is to add Options in the check
	if request.Method != http.MethodPost && request.Method != http.MethodOptions {
		w.Header().Set("Status-Code", "1")
		w.Header().Set("Status-Text", "Method Not Allowed should be POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if request.Body == nil {
		w.Header().Set("Status-Code", "3")
		w.Header().Set("Status-Text", "Body expected")
		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	t := new(Test)

	
	err := json.NewDecoder(request.Body).Decode(&t)
	
	if err != nil {
		w.Header().Set("Status-Code", "3")
		w.Header().Set("Status-Text", "Body expected")
		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	// IMPORTANT: Must close 
	defer request.Body.Close()

	fmt.Printf("TestPost: We got a POST REQUEST WOOOO HOOO!!! Test data: %v\n", t.Text)
}

// TestGet will return some test data
func (r Routes) TestGet(w http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet && request.Method != http.MethodOptions {
		w.Header().Set("Status-Code", "2")
		w.Header().Set("Status-Text", "Method Not Allowed should be GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	t := new(Test)
	t.Text = "Wazaaaaa"

	tMarshalObject, err := json.Marshal(t)

	if err != nil {
		w.Header().Set("Status-Code", "4")
		w.Header().Set("Status-Text", "Ups, I did it again...")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(tMarshalObject))

	return
}