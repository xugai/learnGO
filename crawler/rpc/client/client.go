package main

import (
	"fmt"
	rpcdemo "learnGO/crawler/rpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {

	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{Arg1: 3, Arg2: 4}, &result)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", result)
	}

	err = client.Call("DemoService.Div", rpcdemo.Args{Arg1: 10, Arg2: 0}, &result)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", result)
	}
}
