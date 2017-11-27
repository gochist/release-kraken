package main

import (
	"log"
	"net/http"
)

type Hello struct{}

func main() {
	log.Fatal(
		http.ListenAndServe(":4000", http.FileServer(http.Dir("./www"))))
}
