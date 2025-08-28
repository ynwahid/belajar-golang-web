package main

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Executing Handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Executing Handler")
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Foo Executed")
		fmt.Fprint(writer, "Hello Foo")
	})
	logMiddleware := &LogMiddleware{}
	logMiddleware.Handler = mux

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: logMiddleware,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
