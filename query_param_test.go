package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Ucup", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Body
	body, _ := io.ReadAll(response)
	fmt.Println(string(body))
}

func MultipleQueryParameters(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")
	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParams(t *testing.T) {
	request := httptest.NewRequest(
		http.MethodGet,
		"http://localhost:8080/hello?first_name=Ucup&last_name=Wahid",
		nil,
	)
	recorder := httptest.NewRecorder()

	MultipleQueryParameters(recorder, request)

	response := recorder.Body
	body, _ := io.ReadAll(response)
	fmt.Println(string(body))
}
