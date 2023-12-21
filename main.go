package main

import (
	"fmt"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go file"))
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		fmt.Println(err)
	}
}
