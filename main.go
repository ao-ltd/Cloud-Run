package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Cloud Run では PORT 環境変数が渡される
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Cloud Run with Go!")
	})

	log.Printf("TEST")
	log.Printf("TEST2")
	log.Printf("TEST3")
	log.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
