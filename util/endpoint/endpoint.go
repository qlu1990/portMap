package endpoint

import "strconv"

type Transport int

// const (
// 	TcpProtocol Transport = iota + 1
// 	UdpProtocol
// )

type EndpointAddr struct {
	Ip        string
	Port      int
	Transport string
}

func (endpoint *EndpointAddr) Network() string {
	return endpoint.Transport
}

func (endpoint *EndpointAddr) String() string {
	return endpoint.Ip + ":" + strconv.Itoa(endpoint.Port)
}
