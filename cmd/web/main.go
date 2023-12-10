package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static",fileServer))

	err := http.ListenAndServe(":10000", mux)
	log.Fatal(err)
}