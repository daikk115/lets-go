package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

// https://golang.google.cn/pkg/net/rpc/
// Only methods that satisfy these criteria will be made available for remote access; other methods will be ignored:
// 1. the method's type (Receiver) is exported.
// 2. the method is exported.
// 3. the method has two arguments, both arguments types are exported or must be builtin types.
// 4. the method's second argument is a pointer.
// 5. the method has return type error.
// func (t *T) MethodName(argType T1, replyType *T2) error

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}

func (p *HelloService) XinChao(request string, reply *string) error {
	*reply = "Xin chào " + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Listen TCP error:", err)
	}

	log.Println("Server is ready")

	for {
		// accept connection
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		// serve client in another goroutine
		go func() {
			log.Println("Accept new client:", conn.RemoteAddr())
			rpc.ServeConn(conn)
		}()
	}
}
