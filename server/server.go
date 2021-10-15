package raft_server

import (
	"encoding/json"
	"fmt"

	"github.com/abhishekkr/gol/golnet"
	"github.com/hashicorp/raft"

	raft_cfg "github.com/abhishekkr/rafter/cfg"
	raft_model "github.com/abhishekkr/rafter/model"
	raft_node "github.com/abhishekkr/rafter/node"
	raft_util "github.com/abhishekkr/rafter/util"
)

var (
	ThisNode *raft_node.Node
)

func New() {
	ThisNode = raft_node.New(&raft.Raft{})
	go golnet.TCPServer(
		fmt.Sprintf("%s:%s", raft_cfg.Address, raft_cfg.Port),
		tcpHandler,
	)
}

func errBody(body string, err error) []byte {
	response := raft_model.ServerResponse{
		Body: []byte(body),
		Err:  err,
	}
	blob, errGlob := raft_util.Gob(response)
	if errGlob != nil {
		return []byte(errGlob.Error())
	}
	return blob
}

func tcpHandler(reqBytes []byte) []byte {
	req := raft_model.RpcRequest{}
	if errUnmarshall := json.Unmarshal(reqBytes, &req); errUnmarshall != nil {
		return errBody("raft:server malformed request", errUnmarshall)
	}
	switch req.Action {
	case "stats":
		return NodeStats()
	case "join":
		return NodeJoin()
	case "leave":
		return NodeLeave()
	case "payload":
		return Payload()
	default:
		errAction := fmt.Errorf("unidentified request: %v", req)
		return errBody("raft:server unhandled action", errAction)
	}
}

func NodeStats() []byte {
	blob, errGlob := raft_util.Gob(ThisNode.Stats())
	if errGlob != nil {
		return []byte(errGlob.Error())
	}
	return blob
}

func NodeJoin() []byte {
	req := raft_node.NodeMembershipRequest{
		NodeID:      raft_cfg.ID,
		RaftAddress: raft_cfg.Address,
	}
	blob, errGlob := raft_util.Gob(ThisNode.JoinRequest(req))
	if errGlob != nil {
		return []byte(errGlob.Error())
	}
	return blob
}

func NodeLeave() []byte {
	req := raft_node.NodeMembershipRequest{
		NodeID: raft_cfg.ID,
	}
	blob, errGlob := raft_util.Gob(ThisNode.RemoveRequest(req))
	if errGlob != nil {
		return []byte(errGlob.Error())
	}
	return blob
}

func Payload() []byte {
	req := raft_node.NodeMembershipRequest{
		NodeID: raft_cfg.ID,
	}
	blob, errGlob := raft_util.Gob(ThisNode.RemoveRequest(req))
	if errGlob != nil {
		return []byte(errGlob.Error())
	}
	return blob
}
