package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))

	http.Handle("/", fs)
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	log.Println("Listening...")

	http.ListenAndServe(":3000", nil)
}
