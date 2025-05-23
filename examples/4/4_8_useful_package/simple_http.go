package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	r, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	b, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err == nil {
		fmt.Printf("%s", string(b))
	}
}
