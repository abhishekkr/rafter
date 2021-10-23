package main

import (
	"flag"
	"fmt"

	raft_cfg "github.com/abhishekkr/rafter/cfg"
	raft_client "github.com/abhishekkr/rafter/client"
	raft_model "github.com/abhishekkr/rafter/model"
)

var (
	targetConnectionString = flag.String("target", "", "Connection string for target TCP server as HOST:PORT (127.0.0.1:6661)")
	action                 = flag.String("action", "stats", "Client action string")
	nodeID                 = flag.String("id", "", "Node ID for Membership Request")
	nodeAddress            = flag.String("address", "", "Node Address for Membership Request")
	payloadFile            = flag.String("payload-file", "", "File with Payload content to be sent")
)

func main() {
	if *targetConnectionString == "" {
		*targetConnectionString = fmt.Sprintf("%s:%s", raft_cfg.TargetAddress, raft_cfg.TargetPort)
	}
	fmt.Printf(`RAFTER client..
	Target TCP Server at Port %s
	`, *targetConnectionString)

	client := raft_client.New(*targetConnectionString)
	act(&client)
}

func act(client *raft_client.TCPClient) {
	flag.Parse()
	switch *action {
	case raft_model.RequestActions["stats"]:
		fmt.Printf("%v", client.NodeStats())
	case raft_model.RequestActions["join"]:
		client.NodeJoin(*nodeID, *nodeAddress)
	case raft_model.RequestActions["leave"]:
		client.NodeLeave(*nodeID)
	case raft_model.RequestActions["payload"]:
		fmt.Printf("WIP to send payload from: %s", *payloadFile)
	}
}
