package main

import (
	"fmt"
	"net/http"
	"os"
)

var server = http.NewServeMux()

func main() {
	server.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/", RequestHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	}
	fmt.Printf("Listening on %s\n", port)
	error := http.ListenAndServe(port, nil)
    if error != nil {
        fmt.Printf("Could not start server - try adding a : to your port\n")
    }
}

func RequestHandler(writer http.ResponseWriter, response *http.Request) {
	server.ServeHTTP(writer, response)
}
