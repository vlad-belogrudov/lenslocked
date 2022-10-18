package main

import (
	"fmt"
	"net/http"
	"time"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerFunc)
	server := &http.Server{
		Addr:         ":3000",
		Handler:      mux,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	fmt.Println("Starting the server on :3000...")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
