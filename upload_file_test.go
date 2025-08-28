package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// const maxSize = 100 << 20

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "upload-form.gohtml", nil)
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	// request.ParseMultipartForm(maxSize)
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}

	finalDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(finalDestination, file)
	if err != nil {
		panic(err)
	}

	name := request.PostFormValue("name")
	myTemplates.ExecuteTemplate(writer, "upload-successful.gohtml", map[string]any{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/globe.svg
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Ucup Nur Wahid")
	file, _ := writer.CreateFormFile("file", "CONTOHUPLOAD.svg")
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	respBody, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(respBody))
}
