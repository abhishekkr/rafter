package raft_server

import (
	raft_cfg "github.com/abhishekkr/rafter/cfg"
	raft_client "github.com/abhishekkr/rafter/client"
	raft_model "github.com/abhishekkr/rafter/model"
	raft_o11y "github.com/abhishekkr/rafter/o11y"
	raft_util "github.com/abhishekkr/rafter/util"
)

func (svr *TCPServer) Stats() []byte {
	raft_o11y.Log(svr.nodeHandler.Stats())
	blob, errGlob := raft_util.Gob(svr.nodeHandler.Stats())
	if errGlob != nil {
		return []byte(errGlob.Error())
	}
	return blob
}

func (svr *TCPServer) Join(reqBody []byte) []byte {
	req := raft_model.NodeMembershipRequest{}
	if errUnmarshal := req.Unmarshal(reqBody); errUnmarshal != nil {
		return []byte(errUnmarshal.Error())
	}

	blob, errGlob := raft_util.Gob(svr.nodeHandler.JoinRequest(req))
	if errGlob != nil {
		return []byte(errGlob.Error())
	}
	return blob
}

func (svr *TCPServer) Leave(reqBody []byte) []byte {
	req := raft_model.NodeMembershipRequest{}
	if errUnmarshal := req.Unmarshal(reqBody); errUnmarshal != nil {
		return []byte(errUnmarshal.Error())
	}

	blob, errGlob := raft_util.Gob(svr.nodeHandler.RemoveRequest(req))
	if errGlob != nil {
		return []byte(errGlob.Error())
	}
	return blob
}

func (svr *TCPServer) Payload(reqBody []byte) []byte {
	blob, errGlob := raft_util.Gob(svr.nodeHandler.PayloadHandler(reqBody))
	if errGlob != nil {
		return []byte(errGlob.Error())
	}
	return blob
}

func (svr *TCPServer) SendJoin(sendToBody []byte) []byte {
	var sendTo raft_model.SendTo
	if errUnmarshal := sendTo.Unmarshal(sendToBody); errUnmarshal != nil {
		return []byte(errUnmarshal.Error())
	}

	client := raft_client.New(sendTo.ConnectionString)
	if response := client.NodeJoin(raft_cfg.ID, raft_cfg.Address); response.Error != nil {
		return []byte(response.Error.Error())
	}
	return []byte(raft_cfg.SuccessCode)
}

func (svr *TCPServer) SendLeave(sendToBody []byte) []byte {
	var sendTo raft_model.SendTo
	if errUnmarshal := sendTo.Unmarshal(sendToBody); errUnmarshal != nil {
		return []byte(errUnmarshal.Error())
	}

	client := raft_client.New(sendTo.ConnectionString)
	if response := client.NodeLeave(raft_cfg.ID); response.Error != nil {
		return []byte(response.Error.Error())
	}
	return []byte(raft_cfg.SuccessCode)
}

func (svr *TCPServer) SendPayload(sendToBody []byte) []byte {
	var sendTo raft_model.SendTo
	if errUnmarshal := sendTo.Unmarshal(sendToBody); errUnmarshal != nil {
		return []byte(errUnmarshal.Error())
	}

	client := raft_client.New(sendTo.ConnectionString)
	if response := client.Payload(sendTo.Payload); response.Error != nil {
		return []byte(response.Error.Error())
	}
	return []byte(raft_cfg.SuccessCode)
}
