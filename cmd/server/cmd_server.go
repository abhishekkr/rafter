package main

import (
	"fmt"

	raft_cfg "github.com/abhishekkr/rafter/cfg"
	raft_server "github.com/abhishekkr/rafter/server"
)

func main() {
	fmt.Printf(`
	RAFTER Server up..
	Raft Node (%s) ready at %s:%s
	TCP Server listening at Port %s
	`, raft_cfg.ID, raft_cfg.Address, raft_cfg.Port, raft_cfg.ServerPort)

	raft_server.New()
	fmt.Println("bye.")
}
