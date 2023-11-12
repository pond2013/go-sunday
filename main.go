package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func FuncIndex(w http.ResponseWriter, r *http.Request) {
	// check query param "name"
	fmt.Println(r.URL.Query().Get("name"))

	// check method
	fmt.Println(r.Method)
	if r.Method == "GET" {
		io.WriteString(w, "Hello, World!. This is GET method")
	} else {
		io.WriteString(w, "Hello, World!. This is Other method")
	}
}

func main() {
	http.HandleFunc("/", FuncIndex)
	fmt.Println("start server completed")
	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		log.Fatal(error)
	}
}
