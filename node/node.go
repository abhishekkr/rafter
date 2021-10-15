package raft_node

import (
	"fmt"

	"github.com/hashicorp/raft"
)

var (
	raftError = map[string]int{
		"add-voter":  422,
		"not-leader": 422,
		"bad-node":   500,
	}
)

type Node struct {
	r *raft.Raft
}

type NodeMembershipRequest struct {
	NodeID      string `json:"node_id"`
	RaftAddress string `json:"raft_address"`
}

func New(r *raft.Raft) *Node {
	return &Node{r: r}
}

func (n *Node) Stats() map[string]interface{} {
	return map[string]interface{}{
		"message": "Raft::Stats",
		"data":    n.r.Stats(),
	}
}

func (n *Node) JoinRequest(req NodeMembershipRequest) map[string]interface{} {
	if n.r.State() != raft.Leader {
		return map[string]interface{}{
			"error":    "join request sent to a non-leader node",
			"err-code": raftError["not-leader"],
		}
	}

	if cfgFuture := n.r.GetConfiguration(); cfgFuture.Error() != nil {
		return map[string]interface{}{
			"error":    fmt.Sprintf("raft node-join error with failure to get raft config: %s", cfgFuture.Error()),
			"err-code": raftError["bad-node"],
		}
	}

	voterID := raft.ServerID(req.NodeID)
	voterAddress := raft.ServerAddress(req.RaftAddress)
	if futureNode := n.r.AddVoter(voterID, voterAddress, 0, 0); futureNode.Error() != nil {
		return map[string]interface{}{
			"error":    fmt.Sprintf("raft node-join error with failure to add voter: %s", futureNode.Error()),
			"err-code": raftError["add-voter"],
		}
	}

	return map[string]interface{}{
		"message": fmt.Sprintf("Raft::Join new voter: %s at %s", req.NodeID, req.RaftAddress),
		"data":    n.r.Stats(),
	}
}

func (n *Node) RemoveRequest(req NodeMembershipRequest) map[string]interface{} {
	if n.r.State() != raft.Leader {
		return map[string]interface{}{
			"error":    "remove request sent to a non-leader node",
			"err-code": raftError["not-leader"],
		}
	}

	if cfgFuture := n.r.GetConfiguration(); cfgFuture.Error() != nil {
		return map[string]interface{}{
			"error":    fmt.Sprintf("raft node-remove error with failure to get raft config: %s", cfgFuture.Error()),
			"err-code": raftError["bad-node"],
		}
	}

	voterID := raft.ServerID(req.NodeID)
	if futureNode := n.r.RemoveServer(voterID, 0, 0); futureNode.Error() != nil {
		return map[string]interface{}{
			"error":    fmt.Sprintf("raft node-remove error with failure to remove server: %s", futureNode.Error()),
			"err-code": raftError["add-voter"],
		}
	}

	return map[string]interface{}{
		"message": fmt.Sprintf("Raft::Remove existing node: %s", req.NodeID),
		"data":    n.r.Stats(),
	}
}

func (n *Node) Payload(body []byte) map[string]interface{} {
	if n.r.State() != raft.Leader {
		return map[string]interface{}{
			"error":    "payload request sent to a non-leader node",
			"err-code": raftError["not-leader"],
		}
	}

	return map[string]interface{}{
		"message": fmt.Sprintf("Raft::Payload from node: %s", n.r.String()),
		"data":    body,
	}
}
