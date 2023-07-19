package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello, World! Served from: %s", hostname)
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":8080", nil)
}

