package main

import (
    "fmt"
    "net/http"
	"os"
	"log"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func main() {
	if os.Getenv("PORT") == "" {
		// Default to 8080
		os.Setenv("PORT", "8080")
	}
    http.HandleFunc("/hello", hello)
	log.Println("Listening on port " + os.Getenv("PORT"))
    http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}