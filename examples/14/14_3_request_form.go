package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		name := r.FormValue("name")
		phone := r.FormValue("phone")
		fmt.Fprintf(w, "We confirmed ticket for %s with phone number %s!", name, phone)
	} else {
		tmpl := template.Must(template.ParseFiles("templates/template_form.html"))
		tmpl.Execute(w, "Nothing")
	}
}

func main() {
	http.HandleFunc("/", formHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
