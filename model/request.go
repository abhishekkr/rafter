package raft_model

import (
	raft_util "github.com/abhishekkr/rafter/util"
)

type Client interface {
	Request([]byte) []byte
}

type RpcRequest struct {
	ID     string
	Action string
	Body   []byte
	Err    error
}

func (req *RpcRequest) Send(client Client) ServerResponse {
	response := ServerResponse{}
	reqGob, err := raft_util.Gob(req)
	if err != nil {
		req.Err = err
		return response
	}

	bufResponse := client.Request(reqGob)
	req.Err = raft_util.UnGob(bufResponse, response)
	return response
}
