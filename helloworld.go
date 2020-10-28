package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type returnTest struct {
	Output string `json:"hello"`
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/test_json", jsonHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/test_json" {
		http.NotFound(w, r)
		return
	}

	test := returnTest{
		Output: "World",
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(test)
	// fmt.Fprint(w, "Hello, World!")
}
