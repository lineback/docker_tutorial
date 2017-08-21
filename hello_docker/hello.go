package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello, %v", name)
}

func main() {
	http.HandleFunc("/greeter", sayHello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
