package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/hyprhex/museum/data"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go file"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.templ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	
	html.Execute(w, data.GetAll()[1])
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTemplate)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":8888", server)
	if err != nil {
		fmt.Println(err)
	}
}
