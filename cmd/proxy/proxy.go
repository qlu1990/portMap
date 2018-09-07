package main

import (
	"log"
	"manager/proxy"
	"net"

	"net/http"
	"net/rpc"
)

func main() {
	netproxy := proxy.NewProxy()
	rpc.Register(netproxy)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalf("proxy start listening port 1234 error:%v", err)
	}
	log.Printf("proxy running and listening port:%d", 1234)
	http.Serve(l, nil)

}
