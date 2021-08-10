package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello from go server\n")
}

func main() {
	http.HandleFunc("/hello", hello)
	log.Println("INFO", "server is starting on 8080 port")
	http.ListenAndServe(":8000", nil)
}
