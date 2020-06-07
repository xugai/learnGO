package rpcsupport

import (
	"learnGO/crawler_distributed/config"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRPC(host string, service interface{}) error {
	err := rpc.Register(service)
	if err != nil {
		return err
	}

	listener, err := net.Listen(config.SERVICE_PROTOCOL, host)
	log.Printf("listening %s\n", host)
	if err != nil {
		return err
	}

	for {
		accept, err := listener.Accept()
		if err != nil {
			log.Printf("establish connection error %v\n", err)
			continue
		}
		go jsonrpc.ServeConn(accept)
	}
	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial(config.SERVICE_PROTOCOL, host)
	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil
}
