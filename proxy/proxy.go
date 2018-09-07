package proxy

import (
	"context"
	"log"
	"manager/util/endpoint"
	"manager/util/iptables"
	"net"
	"strconv"
)

type Proxy struct {
	runner      *iptables.Iptables
	backContext context.Context
	cancelCache map[int]context.CancelFunc
}
type Args struct {
	Hostport int
	Endpoint endpoint.EndpointAddr
}
type ListenArgs struct {
	Transport string
	Port      int
}
type MapError struct {
	Msg string
}

var FuncMap map[string]string = map[string]string{"D": "Proxy.DeleteMapPort", "A": "Proxy.AddMapPort"}

func (err *MapError) Error() string {
	return err.Msg
}

type Reply struct {
	Status int
	Error  string
}

func NewProxy() *Proxy {
	netproxy := Proxy{
		runner:      iptables.New(),
		backContext: context.Background(),
		cancelCache: make(map[int]context.CancelFunc),
	}
	return &netproxy
}

func (netproxy *Proxy) AddMapPort(args *Args, reply *Reply) error {
	if _, ok := netproxy.cancelCache[args.Hostport]; ok {
		reply.Status = -2
		reply.Error = "the port is exits"
	}
	ctx, cancel := context.WithCancel(netproxy.backContext)
	listenArgs := ListenArgs{
		Port:      args.Hostport,
		Transport: args.Endpoint.Network(),
	}
	go ListenPort(ctx, listenArgs)
	netproxy.cancelCache[args.Hostport] = cancel
	var table iptables.Table = "nat"
	var chain iptables.Chain = "PREROUTING"
	cmdArgs := []string{"-p", args.Endpoint.Network(), "--dport", strconv.Itoa(args.Hostport), "-j", "DNAT",
		"--to-destination", args.Endpoint.String()}
	if exit, _ := netproxy.runner.EnsureRule(table, chain, cmdArgs...); !exit {
		err := netproxy.runner.AddRule(table, chain, cmdArgs...)
		if err != nil {
			log.Printf("Error in AddMapPort rule:%v", err)
			reply.Status = -1
			reply.Error = err.Error()

		}
		reply.Status = 0
		reply.Error = ""
		return err
	} else {
		reply.Status = -2
		reply.Error = "the port is exits"
		return nil
	}

}

func (netproxy *Proxy) DeleteMapPort(args *Args, reply *Reply) error {
	if calcel, ok := netproxy.cancelCache[args.Hostport]; !ok {
		reply.Status = -3
		reply.Error = "the port is not exits"
	} else {
		calcel()
		delete(netproxy.cancelCache, args.Hostport)
		// log.Println("calcel listens", calcel)
	}

	var table iptables.Table = "nat"
	var chain iptables.Chain = "PREROUTING"
	err := netproxy.runner.DeleteRule(table, chain, "-p", args.Endpoint.Network(), "--dport", strconv.Itoa(args.Hostport), "-j", "DNAT",
		"--to-destination", args.Endpoint.String())
	if err != nil {
		log.Printf("Error in DeleteMapPort rule:%v", err)
		reply.Status = -4
		reply.Error = err.Error()

	}
	reply.Status = 0
	reply.Error = ""
	return err

}

func ListenPort(ctx context.Context, args ListenArgs) {
	listener, err := net.Listen(args.Transport, ":"+strconv.Itoa(args.Port))
	if err != nil {
		log.Printf("listen error port:%d ,error:%v", args.Port, err)
	}

	// listener.Accept()

	select {
	case <-ctx.Done():
		log.Printf("stop listen port:%d", args.Port)
		listener.Close()
		return
	}
}
