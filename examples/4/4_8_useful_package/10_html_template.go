package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("4_8_useful_package/template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := User{Name: "Go Developer"}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
