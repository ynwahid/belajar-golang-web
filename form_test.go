package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func PostForm(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	// Shortcut
	// request.PostFormValue("first_name")
	// request.PostFormValue("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestPostForm(t *testing.T) {
	requestBody := strings.NewReader("first_name=Ucup&last_name=Wahid")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	PostForm(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
