package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "jsonplaceholder.typicode.com:443")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Connected to server. Type messages:")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		// Send text to server
		fmt.Fprintf(conn, text+"\n")

		// Read response from server
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			break
		}

		fmt.Print("Server says: " + response)
	}
}
