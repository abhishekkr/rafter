package raft_client

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/rogpeppe/fastuuid"

	raft_cfg "github.com/abhishekkr/rafter/cfg"
	raft_controller "github.com/abhishekkr/rafter/controller"
	raft_model "github.com/abhishekkr/rafter/model"
	raft_util "github.com/abhishekkr/rafter/util"
)

var (
	uuid = fastuuid.MustNewGenerator()
)

type TCPClient struct {
	Connection net.Conn
}

func createTCPClient(connectionString string) TCPClient {
	connection, err := net.Dial("tcp", connectionString)
	if err != nil {
		fmt.Println("ERROR: Golnet TCP client failed to connect to server.", err)
		os.Exit(1)
	}
	return TCPClient{
		Connection: connection,
	}
}

func (client TCPClient) Send(req raft_model.RpcRequest) raft_model.ServerResponse {
	response := raft_model.ServerResponse{}
	reqGob, err := raft_util.Gob(req)
	if err != nil {
		req.Err = err
		return response
	}

	_, write_err := client.Connection.Write(reqGob)
	if write_err != nil {
		fmt.Println("ERROR: Golnet TCP client failed to write to server.", write_err)
	}
	client.Connection.(*net.TCPConn).CloseWrite()

	readBuffer, read_err := ioutil.ReadAll(client.Connection)
	if read_err != nil {
		fmt.Println("ERROR: Golnet TCP client failed to read from server.", read_err)
	}
	req.Err = raft_util.UnGob(readBuffer, &response)
	return response
}

func (client TCPClient) NodeStats() raft_model.ServerResponse {
	req := raft_model.RpcRequest{
		ID:     fmt.Sprintf("%s:STATS:%s", raft_cfg.ID, uuid.Hex128()),
		Action: raft_model.RequestActions["stats"],
	}
	response := client.Send(req)
	if req.Err != nil && response.Error == nil {
		response.ErrCode = raft_controller.RaftError["bad-request"]
		response.Error = req.Err
	}
	return response
}

func (client TCPClient) NodeJoin(nodeID string, nodeAddress string) raft_model.ServerResponse {
	body, errGob := raft_util.Gob(raft_model.NodeMembershipRequest{
		NodeID:      nodeID,
		RaftAddress: nodeAddress,
	})
	if errGob != nil {
		return raft_model.ServerResponse{
			ErrCode: raft_controller.RaftError["bad-request"],
			Error:   errGob,
		}
	}
	req := raft_model.RpcRequest{
		ID:     fmt.Sprintf("%s:JOIN:%s", raft_cfg.ID, uuid.Hex128()),
		Action: raft_model.RequestActions["join"],
		Body:   body,
	}
	response := client.Send(req)
	if req.Err != nil && response.Error == nil {
		response.ErrCode = raft_controller.RaftError["bad-request"]
		response.Error = req.Err
	}
	return response
}

func (client TCPClient) NodeLeave(nodeID string) raft_model.ServerResponse {
	body, errGob := raft_util.Gob(raft_model.NodeMembershipRequest{
		NodeID: nodeID,
	})
	if errGob != nil {
		return raft_model.ServerResponse{
			ErrCode: raft_controller.RaftError["bad-request"],
			Error:   errGob,
		}
	}
	req := raft_model.RpcRequest{
		ID:     fmt.Sprintf("%s:LEAVE:%s", raft_cfg.ID, uuid.Hex128()),
		Action: raft_model.RequestActions["leave"],
		Body:   body,
	}
	response := client.Send(req)
	if req.Err != nil && response.Error == nil {
		response.ErrCode = raft_controller.RaftError["bad-request"]
		response.Error = req.Err
	}
	return response
}

func (client TCPClient) Payload(body []byte) raft_model.ServerResponse {
	req := raft_model.RpcRequest{
		ID:     fmt.Sprintf("%s:PAYLOAD:%s", raft_cfg.ID, uuid.Hex128()),
		Action: raft_model.RequestActions["payload"],
		Body:   body,
	}
	response := client.Send(req)
	if req.Err != nil && response.Error == nil {
		response.ErrCode = raft_controller.RaftError["bad-request"]
		response.Error = req.Err
	}
	return response
}
