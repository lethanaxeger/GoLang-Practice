package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Halaman Utama")
	})

	http.HandleFunc("/Profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Halaman profile")
	})
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
	// fmt.Println("Hello")
}
