package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	fsh := http.FileServer(http.Dir("/usr/share/doc"))
	log.Fatal(http.ListenAndServe(":8080", fsh))
}
