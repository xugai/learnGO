package main

import (
	rpcdemo "learnGO/crawler/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	err := rpc.Register(rpcdemo.DemoService{})
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		accept, err := listener.Accept()
		if err != nil {
			log.Printf("establish connection error %v\n", err)
			continue
		}
		go jsonrpc.ServeConn(accept)
	}
}
