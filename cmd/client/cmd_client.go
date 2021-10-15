package main

import (
	"fmt"

	raft_cfg "github.com/abhishekkr/rafter/cfg"
	raft_rpc "github.com/abhishekkr/rafter/rpc"
)

var (
	action = "stats"
)

func main() {
	fmt.Printf(`RAFTER client..
	Target TCP Server at Port %s:%s
	`, raft_cfg.TargetAddress, raft_cfg.TargetPort)
	if action != "stats" {
		fmt.Printf(`RAFTER client..
	Source Raft Node (%s) at %s:%s & %s
	`, raft_cfg.ID, raft_cfg.Address, raft_cfg.Port, raft_cfg.ServerPort)
	}

	client := raft_rpc.New()
	fmt.Printf("%v", client)
}
