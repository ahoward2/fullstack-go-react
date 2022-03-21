package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	if r.URL.Path != "/hello" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
	}

	if r.Method != "GET" {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
	}

	type HelloWorld struct {
		Message string `json:"message"`
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(HelloWorld{Message: "Hello World"})
 
}

func main() {

	// Server static react build (don't have to serve the static files if you don't want to,
  // you could instead just use the json endpoints).
	fileServer := http.FileServer(http.Dir("./dist"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	

	fmt.Printf("Starting server at port 8080 \n")	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}