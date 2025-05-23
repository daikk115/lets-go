package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:]) // Eg: /Alice → Hello, Alice!
}

func main() {
	http.HandleFunc("/", helloHandler) // mapping handler with /
	fmt.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

// Example curl
// curl http://localhost:8080/daidv
