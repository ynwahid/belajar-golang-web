package main

import (
	"fmt"
	"net/http"
	"testing"
)

type ErrorMiddleware struct {
	Handler http.Handler
}

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *ErrorMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error: %v", err)
		}
	}()
	middleware.Handler.ServeHTTP(writer, request)
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
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Panic Executed")
		panic("a panic occurred in the system")
	})
	logMiddleware := &LogMiddleware{}
	logMiddleware.Handler = mux

	errorMiddleware := &ErrorMiddleware{}
	errorMiddleware.Handler = logMiddleware

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorMiddleware,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
