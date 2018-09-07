package node

import (
	"context"
	"manager/node/control"
	"net"
	"sync"
)

// func init() {

// }

// func main() {
// 	var cmdString iptables.Cmd = "/bin/ls"
// 	iptables.IptablesCmd(cmdString)

// }
type node struct {
	ServerAddr  net.Addr
	ProxyPort   int
	CancelCache map[string]context.CancelFunc
}

func New(servcerAddr net.Addr, proxyPort int) *node {
	nodeRunner := &node{
		ServerAddr: servcerAddr,
		ProxyPort:  proxyPort,
	}
	return nodeRunner

}
func (nodeRunner *node) Run() {
	bctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(1)
	proxyControl := control.NewProxy()
	ctx1, calcel := context.WithCancel(bctx)
	nodeRunner.CancelCache["proxy"] = calcel
	go proxyControl.Run(ctx1, "127.0.0.1", nodeRunner.ProxyPort)

}
