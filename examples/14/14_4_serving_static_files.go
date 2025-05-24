package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Return a simple HTML page with a link to a static file
	html := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Hello</title>
    </head>
    <body>
        <h1>Hello, World!</h1>
        <p>View the static file: <a href="/static/" target="_blank">Click here</a></p>
    </body>
    </html>
    `
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.HandleFunc("/", helloHandler)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
