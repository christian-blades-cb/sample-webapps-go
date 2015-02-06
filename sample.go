package main

import (
	"fmt"
	"net/http"
	"os"
)

func biggerHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/sample_text.txt")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func smallsHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello. I am %s", hostname)
}

func main() {
	http.HandleFunc("/bigger", biggerHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/", smallsHandler)
	http.ListenAndServe(":8080", nil)
}
