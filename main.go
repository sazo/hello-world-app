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
		fmt.Fprintf(w, "Hallo world! From host: %s", hostname)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
