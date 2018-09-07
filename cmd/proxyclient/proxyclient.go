package main

import (
	"fmt"
	"log"
	"manager/proxy"
	"manager/util/endpoint"
	"net/rpc"
	"os"
	"strconv"
)

func main() {

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	if len(os.Args) < 6 {
		fmt.Println("arg err")
		fmt.Println("useage: proxyclient tcp 3306 172.24.8.80 3306 A")
		fmt.Println("arg1 协议")
		fmt.Println("arg2 hostport")
		fmt.Println("arg3 destiantionhost")
		fmt.Println("arg4 destinationport")
		fmt.Println("arg5 A is add and  D is delete ")
		return
	}
	prot := os.Args[1]
	hostport, err := strconv.Atoi(os.Args[2])
	desthost := os.Args[3]
	destport, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Println("port must a number")
		return
	}
	test_endpoint := endpoint.EndpointAddr{
		Ip:        desthost,
		Port:      destport,
		Transport: prot,
	}
	args := &proxy.Args{
		Hostport: hostport,
		Endpoint: test_endpoint,
	}
	var reply proxy.Reply
	if len(os.Args) >= 6 {

		log.Println(os.Args)
		if os.Args[5] == "D" {
			log.Println("delete rule")
			err = client.Call("Proxy.DeleteMapPort", args, &reply)
		} else if os.Args[5] == "A" {
			log.Println("add rule")
			err = client.Call("Proxy.AddMapPort", args, &reply)
		} else {
			return
		}
		if err != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("status: %d,error:%s\n", reply.Status, reply.Error)
	} else {
		return
	}

}
