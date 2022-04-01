package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

//3. write tests for service layer

func Mux() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/api", apiHandler)
	return router
}

func recordRequest(mux *http.ServeMux, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
	//assert.Equal(t, expected, actual, "ERROR") //"Expected status %d. Got %d", expected, actual
}

func TestPost(t *testing.T) {

	mux := Mux()
	jsonStr := []byte(`{"last": 3}`)
	req, err := http.NewRequest(http.MethodPost, "/api", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	res := recordRequest(mux, req)
	//fmt.Println(res.Body)
	checkResponseCode(t, res.Code, http.StatusOK)
}

func TestGet(t *testing.T) {

	mux := Mux()
	jsonStr := []byte(`{"last": 5}`)
	req, err := http.NewRequest(http.MethodGet, "/api", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	res := recordRequest(mux, req)
	checkResponseCode(t, res.Code, http.StatusOK)
}

func TestError1(t *testing.T) {
	mux := Mux()
	jsonStr := []byte(`{"last": -}`)
	req, err := http.NewRequest(http.MethodGet, "/api", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	res := recordRequest(mux, req)
	//fmt.Println(res.Body)
	var resMap = make(map[string]interface{})
	json.NewDecoder(res.Body).Decode(&resMap)
	errNo, _ := resMap["error"].(float64)
	checkResponseCode(t, 1, int(errNo))
}
