package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, "<h1>Hallo world! From host: %s</h1>", hostname)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
