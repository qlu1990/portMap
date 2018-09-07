package proxy

import (
	"fmt"
	"manager/util/endpoint"
	"testing"
)

func Test_AddMapPort(t *testing.T) {

	test_endpoint := endpoint.EndpointAddr{
		Ip:        "172.24.8.80",
		Port:      3306,
		Transport: "tcp",
	}
	args := &Args{
		Hostport: 3306,
		Endpoint: test_endpoint,
	}
	reply := &Reply{}
	netproxy := NewProxy()
	netproxy.AddMapPort(args, reply)

}

func Test_DeleteMapPort(t *testing.T) {
	test_endpoint := endpoint.EndpointAddr{
		Ip:        "172.24.8.80",
		Port:      3306,
		Transport: "tcp",
	}
	args := &Args{
		Hostport: 3306,
		Endpoint: test_endpoint,
	}
	reply := &Reply{}
	netproxy := NewProxy()
	err := netproxy.DeleteMapPort(args, reply)
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println("deletemapport test success")
		t.Log("deletemapport test success")
	}

}
