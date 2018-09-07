package iptables

import (
	"bytes"
	"os/exec"
	"sync"
)

type operation string

type Table string
type Chain string

// type Cmd string

type Iptables struct {
	mu sync.Mutex
}

const (
	opCreateChain operation = "-N"
	opFlushChain  operation = "-F"
	opDeleteChain operation = "-X"
	opAppendRule  operation = "-A"
	opCheckRule   operation = "-C"
	opDeleteRule  operation = "-D"
)

func New() *Iptables {
	runner := &Iptables{}
	return runner
}
func makeFullArgs(table Table, chain Chain, operate operation, args ...string) []string {

	return append([]string{string(operate), string(chain), "-t", string(table)}, args...)
}

func (ipRunner *Iptables) iptablesCmd(cmdString []string) (string, error) {
	cmd := exec.Command("/sbin/iptables", cmdString...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	// fmt.Println(err)
	// fmt.Println(out.String())
	return out.String(), err

}
func (ipRunner *Iptables) AddRule(table Table, chain Chain, args ...string) error {
	fullArgs := makeFullArgs(table, chain, opAppendRule, args...)
	// ipRunner.mu.Lock()
	// defer ipRunner.mu.Unlock()

	_, err := ipRunner.iptablesCmd(fullArgs)
	return err

}

func (ipRunner *Iptables) DeleteRule(table Table, chain Chain, args ...string) error {
	fullArgs := makeFullArgs(table, chain, opDeleteRule, args...)
	// ipRunner.mu.Lock()
	// defer ipRunner.mu.Unlock()
	_, err := ipRunner.iptablesCmd(fullArgs)
	return err

}

func (ipRunner *Iptables) EnsureRule(table Table, chain Chain, args ...string) (bool, error) {

	fullArgs := makeFullArgs(table, chain, opCheckRule, args...)
	out, err := ipRunner.iptablesCmd(fullArgs)

	if err == nil && len(out) < 1 {
		return true, err
	} else {
		return false, err
	}

}
