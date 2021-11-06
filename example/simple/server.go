package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "hello, %s", r.URL.Path)
	})
	log.Fatal(http.ListenAndServe(":9090", nil))
}
