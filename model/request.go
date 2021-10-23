package raft_model

import (
	"fmt"
	"strings"

	raft_util "github.com/abhishekkr/rafter/util"
)

var (
	RequestActions = map[string]string{
		"stats":        "/stats",
		"join":         "/join",
		"leave":        "/leave",
		"payload":      "/payload",
		"send-join":    "/send/join",
		"send-leave":   "/send/leave",
		"send-payload": "/send/payload",
	}
)

type RpcRequest struct {
	ID     string
	Action string
	Body   []byte
	Err    error
}

type NodeMembershipRequest struct {
	NodeID      string `json:"node_id"`
	RaftAddress string `json:"raft_address"`
}

type SendTo struct {
	ConnectionString string
	Payload          []byte
}

func (req *NodeMembershipRequest) Unmarshal(reqBody []byte) error {
	if errUnmarshall := raft_util.UnGob(reqBody, &req); errUnmarshall != nil {
		return errUnmarshall
	}
	return nil
}

func (req *NodeMembershipRequest) Marshal() ([]byte, error) {
	return raft_util.Gob(req)
}

func (n *SendTo) Unmarshal(nodeBody []byte) error {
	if errUnmarshall := raft_util.UnGob(nodeBody, &n); errUnmarshall != nil {
		return errUnmarshall
	}
	if strings.Index(n.ConnectionString, ":") > 1 && len(n.ConnectionString) > 2 {
		return fmt.Errorf("erroroneous Node connection string")
	}
	return nil
}

func (n *SendTo) Marshal() ([]byte, error) {
	return raft_util.Gob(n)
}
