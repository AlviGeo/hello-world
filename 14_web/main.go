package main

import (
	"fmt"
	"net/http"
)

// F di fmt berfungsi utk mengformat receiver
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to the Hell</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/About", about)
	fmt.Println("Server Starting...")
	// nil is null
	http.ListenAndServe(":3000", nil)
}
