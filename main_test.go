package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestZHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ZHandler)

	handler.ServeHTTP(rr, req)

	// Valid status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	fmt.Printf(rr.Body.String())
	if body := rr.Body.String(); body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expected)
	}

}
