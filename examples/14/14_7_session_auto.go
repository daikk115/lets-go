package main

// go get github.com/gorilla/sessions

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func init() {
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,  // thời gian tồn tại (giây)
		HttpOnly: true,  // ngăn JS truy cập
		Secure:   false, // đổi thành true nếu dùng HTTPS
	}
}

func sessionHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "custome-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "ParseForm() error", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		phone := r.FormValue("phone")
		session.Values["name"] = name
		session.Values["phone"] = phone

		// PHẢI lưu session trước khi ghi response
		err := session.Save(r, w)
		if err != nil {
			http.Error(w, "Cannot save session", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "We confirmed ticket for %s with phone number %s!", name, phone)
		return
	}

	name, nameOk := session.Values["name"].(string)
	phone, phoneOk := session.Values["phone"].(string)

	if nameOk && phoneOk {
		fmt.Fprintf(w, "Session found: %s - %s", name, phone)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/template_form.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", sessionHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
