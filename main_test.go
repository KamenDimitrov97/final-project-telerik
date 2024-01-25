package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRoot(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(getRoot)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "This is my website!\n"
	if recorder.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}

func TestGetHello(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(getHello)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Hello, There!\n"
	if recorder.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}
