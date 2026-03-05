package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello new express")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am from Go TypeScript")
}

func main() {

	mux := http.NewServeMux() //create an router

	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("Server is running on port : 3000")

	err := http.ListenAndServe(":3000", mux) // failed to start the server
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
