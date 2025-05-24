package main

// go get github.com/gorilla/sessions

import (
	"encoding/gob"
	"fmt"

	"github.com/gorilla/securecookie"
)

func init() {
	gob.Register(map[interface{}]interface{}{})
}

var hashKey = []byte("super-secret-key") // must match your session setup
var s = securecookie.New(hashKey, nil)

func manualDecodeSession(cookieValue string) {
	sessionValues := make(map[interface{}]interface{})

	err := s.Decode("custome-session", cookieValue, &sessionValues)
	if err != nil {
		fmt.Println("Failed to decode:", err)
		return
	}

	// Access session values
	for k, v := range sessionValues {
		fmt.Printf("%v: %v\n", k, v)
	}
}

func main() {
	manualDecodeSession("MTc0ODA2MDQ0NHxEWDhFQVFMX2dBQUJFQUVRQUFCR180QUFBZ1p6ZEhKcGJtY01CZ0FFYm1GdFpRWnpkSEpwYm1jTUJRQURSR0ZwQm5OMGNtbHVad3dIQUFWd2FHOXVaUVp6ZEhKcGJtY01EQUFLTURNMU9UWTRNekEyTlE9PXy9_Na-W54TYXGUf4n6EB71YRq_f6iGcxcZ2i7h6NuyUg==")
}
