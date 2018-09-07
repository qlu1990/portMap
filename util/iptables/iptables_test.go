package iptables

import (
	"fmt"
	"testing"
)

var (
	runner *Iptables
	args   []string = []string{"-p", "tcp", "--dport", "3306", "-j", "DNAT", "--to-destination", "172.24.8.80:3306"}
	table  Table    = "nat"
	chain  Chain    = "DOCKER"
)

func Test_AddRule(t *testing.T) {
	if exit, _ := runner.EnsureRule(table, chain, args...); !exit {
		runner.AddRule(table, chain, args...)
	}

}

func Test_DeleteRule(t *testing.T) {
	runner.DeleteRule(table, chain, args...)

}

func Test_EnsureRule(t *testing.T) {
	exit, _ := runner.EnsureRule(table, chain, args...)
	fmt.Println(exit)

}
