package raft_controller

import (
	"fmt"

	"github.com/hashicorp/raft"

	raft_model "github.com/abhishekkr/rafter/model"
)

var (
	RaftError = map[string]int{
		"bad-request": 400,
		"add-voter":   422,
		"not-leader":  422,
		"bad-node":    500,
	}
)

type Node struct {
	r *raft.Raft
}

func New(r *raft.Raft) *Node {
	return &Node{r: r}
}

func (n *Node) Stats() raft_model.ServerResponse {
	return raft_model.ServerResponse{
		Message: "Raft::Stats",
		Data:    n.r.Stats(),
	}
}

func (n *Node) JoinRequest(req raft_model.NodeMembershipRequest) raft_model.ServerResponse {
	if n.r.State() != raft.Leader {
		return raft_model.ServerResponse{
			Error:   fmt.Errorf("join request sent to a non-leader node"),
			ErrCode: RaftError["not-leader"],
		}
	}

	if cfgFuture := n.r.GetConfiguration(); cfgFuture.Error() != nil {
		return raft_model.ServerResponse{
			Error:   fmt.Errorf("raft node-join error with failure to get raft config: %s", cfgFuture.Error()),
			ErrCode: RaftError["bad-node"],
		}
	}

	voterID := raft.ServerID(req.NodeID)
	voterAddress := raft.ServerAddress(req.RaftAddress)
	if futureNode := n.r.AddVoter(voterID, voterAddress, 0, 0); futureNode.Error() != nil {
		return raft_model.ServerResponse{
			Error:   fmt.Errorf("raft node-join error with failure to add voter: %s", futureNode.Error()),
			ErrCode: RaftError["add-voter"],
		}
	}

	return raft_model.ServerResponse{
		Message: fmt.Sprintf("Raft::Join new voter: %s at %s", req.NodeID, req.RaftAddress),
		Data:    n.r.Stats(),
	}
}

func (n *Node) RemoveRequest(req raft_model.NodeMembershipRequest) raft_model.ServerResponse {
	if n.r.State() != raft.Leader {
		return raft_model.ServerResponse{
			Error:   fmt.Errorf("remove request sent to a non-leader node"),
			ErrCode: RaftError["not-leader"],
		}
	}

	if cfgFuture := n.r.GetConfiguration(); cfgFuture.Error() != nil {
		return raft_model.ServerResponse{
			Error:   fmt.Errorf("raft node-remove error with failure to get raft config: %s", cfgFuture.Error()),
			ErrCode: RaftError["bad-node"],
		}
	}

	voterID := raft.ServerID(req.NodeID)
	if futureNode := n.r.RemoveServer(voterID, 0, 0); futureNode.Error() != nil {
		return raft_model.ServerResponse{
			Error:   fmt.Errorf("raft node-remove error with failure to remove server: %s", futureNode.Error()),
			ErrCode: RaftError["add-voter"],
		}
	}

	return raft_model.ServerResponse{
		Message: fmt.Sprintf("Raft::Remove existing node: %s", req.NodeID),
		Data:    n.r.Stats(),
	}
}

func (n *Node) PayloadHandler(body []byte) raft_model.ServerResponse {
	if n.r.State() != raft.Leader {
		return raft_model.ServerResponse{
			Error:   fmt.Errorf("payload request sent to a non-leader node"),
			ErrCode: RaftError["not-leader"],
		}
	}

	return raft_model.ServerResponse{
		Message: fmt.Sprintf("Raft::Payload from node: %s", n.r.String()),
		Data:    body,
	}
}
