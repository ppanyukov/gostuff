package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// NewSimpleHandler creates new instance of the blob storage HTTP handler.
func NewSimpleHandler() (*SimpleHandler, error) {
	return &SimpleHandler{}, nil
}

type SimpleHandler struct{}

// ServeHTTP implements the HTTP handler to list contents of the blob store.
func (server *SimpleHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {

	url := request.URL.String()

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, "OK: %s:\n", url)
	fmt.Fprint(w, "----------------\n")
}

func main() {
	var port string

	flag.StringVar(&port, "port", ":8080", "The port to listen on in format ':PORT', e.g ':8080'.")
	flag.Parse()

	blobStorageHandler, err := NewSimpleHandler()
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	fmt.Println("Will listen on port", port)

	s := &http.Server{
		Addr:           port,
		Handler:        blobStorageHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
