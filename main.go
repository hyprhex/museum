package main

import (
	"fmt"
	"net/http"
)

func main() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello from Go file"))
    })
    err := http.ListenAndServe(":3333", nil)
    if err != nil {
        fmt.Println("Error while opening the server")
    }
}
