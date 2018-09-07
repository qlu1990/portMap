package control

import (
	"context"
	"log"
	"manager/proxy"
	"net/rpc"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

type Job struct {
	operate string
	args    *proxy.Args
	reply   *proxy.Reply
}
type proxy_control struct {
	Jobs chan Job
}

func NewProxy() *proxy_control {
	control := &proxy_control{
		Jobs: make(chan Job, 10),
	}
	return control
}

func (control *proxy_control) Run(ctx context.Context, ip string, port int) {
	defer wg.Done()
	client, err := rpc.DialHTTP("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		log.Printf("rpc http clinet create error:%v", err)
	}

	select {
	case <-ctx.Done():
		return
	case job := <-control.Jobs:
		err := client.Call(proxy.FuncMap[job.operate], job.args, job.reply)
		if err != nil {
			log.Printf("proxy rpc call error func:%s ,error:%v", proxy.FuncMap[job.operate], err)
		}
	}

}
