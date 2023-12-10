package main

import (
	"log"
	"net/http"
)


func home(response http.ResponseWriter, request *http.Request){
	response.Write([]byte("Hello World!"))

}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	err := http.ListenAndServe(":10000", mux)
	log.Fatal(err)

}