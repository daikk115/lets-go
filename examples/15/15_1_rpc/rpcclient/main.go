package main

import (
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Fatal("Dialing error:", err)
	}
	defer client.Close()

	var hello string
	var xinchao string

	if err = client.Call("HelloService.Hello", "David", &hello); err != nil {
		log.Fatal(err)
	}
	if err = client.Call("HelloService.XinChao", "Đại", &xinchao); err != nil {
		log.Fatal(err)
	}

	log.Println(hello)
	log.Println(xinchao)
}
