package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deployed by our pipeline!")
	fmt.Println("API was invoked")
}

func handleRequests() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}

