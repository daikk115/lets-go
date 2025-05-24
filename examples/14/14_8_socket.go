package main

// go get github.com/gorilla/websocket

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)

	// New routine to get msg from client
	go func() {
		for {
			if _, msg, err := conn.ReadMessage(); err != nil {
				// Client close → exit while-loop.
				conn.Close()
				return
			} else {
				fmt.Printf("Received from client: %s\n", msg)
			}
		}
	}()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	i := 0
	// Main routine to sent msg to client
	for {
		i += 1
		time.Sleep(time.Duration(1) * time.Second)
		msg := []byte(fmt.Sprintf("%d", i))
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			fmt.Println("Write error:", err)
			return
		}
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/template_socket.html"))
	tmpl.Execute(w, "Nothing")
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/", formHandler)
	fmt.Println("Socket server started at :8080")
	http.ListenAndServe(":8080", nil)
}
