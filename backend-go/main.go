package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Tambahkan CORS Header
	w.Header().Set("Access-Control-Allow-Origin", "*") // "*" = izinkan semua asal
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Fprint(w, "Hello from Go Backend!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Backend running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
