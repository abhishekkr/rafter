package raft_rpc

import (
	"github.com/abhishekkr/gol/golnet"

	raft_cfg "github.com/abhishekkr/rafter/cfg"
	raft_model "github.com/abhishekkr/rafter/model"
	raft_node "github.com/abhishekkr/rafter/node"
	raft_util "github.com/abhishekkr/rafter/util"
)

func New() raft_model.Client {
	var client *golnet.TCPClient
	// prepare client for TargetAddress:TargetPort
	return client
}

func NodeStats(client raft_model.Client) error {
	req := raft_model.RpcRequest{
		ID:     "${NODE_ID}_STATS_(ABK::put-UUID-here)",
		Action: "stats",
	}
	response := req.Send(client)
	if req.Err != nil {
		return req.Err
	} else if response.Err != nil {
		return response.Err
	}
	return nil
}

func NodeJoin(client raft_model.Client) error {
	body, errGob := raft_util.Gob(raft_node.NodeMembershipRequest{
		NodeID:      raft_cfg.ID,
		RaftAddress: raft_cfg.Address,
	})
	if errGob != nil {
		return errGob
	}
	req := raft_model.RpcRequest{
		ID:     "${NODE_ID}_JOIN_(ABK::put-UUID-here)",
		Action: "join",
		Body:   body,
	}
	response := req.Send(client)
	if req.Err != nil {
		return req.Err
	} else if response.Err != nil {
		return response.Err
	}
	return nil
}

func NodeLeave(client raft_model.Client) error {
	body, errGob := raft_util.Gob(raft_node.NodeMembershipRequest{
		NodeID: raft_cfg.ID,
	})
	if errGob != nil {
		return errGob
	}
	req := raft_model.RpcRequest{
		ID:     "${NODE_ID}_LEAVE_(ABK::put-UUID-here)",
		Action: "leave",
		Body:   body,
	}
	response := req.Send(client)
	if req.Err != nil {
		return req.Err
	} else if response.Err != nil {
		return response.Err
	}
	return nil
}

func Payload(client raft_model.Client, body []byte) error {
	req := raft_model.RpcRequest{
		ID:     "${NODE_ID}_PAYLOAD_(ABK::put-UUID-here)",
		Action: "payload",
		Body:   body,
	}
	response := req.Send(client)
	if req.Err != nil {
		return req.Err
	} else if response.Err != nil {
		return response.Err
	}
	return nil
}
