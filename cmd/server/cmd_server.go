package main

import (
	"fmt"
	"os"

	raft_cfg "github.com/abhishekkr/rafter/cfg"
	raft_server "github.com/abhishekkr/rafter/server"
)

func main() {
	fmt.Printf(`
	RAFTER Server up..
	Raft Node (%s) ready at %s:%s
	TCP Server listening at Port %s
	`, raft_cfg.ID, raft_cfg.Address, raft_cfg.Port, raft_cfg.ServerPort)

	manager := raft_server.New()

	defer func() {
		if err := manager.Store.Close(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error closing FSM Store: %s\n", err.Error())
		}
	}()
	manager.Server.Start(fmt.Sprintf("%s:%s", raft_cfg.Address, raft_cfg.ServerPort))

	fmt.Println("bye.")
}
