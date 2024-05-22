package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enescakir/emoji"
)

func TestHelloHandler(t *testing.T) {
	req,err := http.NewRequest("GET","/",nil)

	if err!=nil{
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloHandler)

	handler.ServeHTTP(rr,req)
	
	if status := rr.Code; status!= http.StatusOK {
		t.Errorf("handler returned wrong status Code: got %v, want: %v",status,http.StatusOK)
	}

	expected := fmt.Sprintf("CICD Testing %v",emoji.Popcorn)

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v",rr.Body.String(),expected)
	}
}