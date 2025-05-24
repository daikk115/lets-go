package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/template.html"))
	tmpl.Execute(w, map[string]string{
		"Title":   "Welcome",
		"Message": fmt.Sprintf("Hello %v from template!", r.URL.Path[1:]),
	})
}

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
