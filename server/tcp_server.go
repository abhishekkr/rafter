package raft_server

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	raft_controller "github.com/abhishekkr/rafter/controller"
	raft_model "github.com/abhishekkr/rafter/model"
	raft_util "github.com/abhishekkr/rafter/util"
)

var (
	TcpServerHalt = false
)

type TCPServer struct {
	nodeHandler *raft_controller.Node
}

func (svr *TCPServer) Start(connection_string string) {
	server, err := net.Listen("tcp", connection_string)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer server.Close()
	for {
		if TcpServerHalt {
			server.Close()
			return
		}
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go svr.handleRequest(connection)
	}
}

func (svr *TCPServer) handleRequest(conn net.Conn) {
	request, read_err := ioutil.ReadAll(conn)
	if read_err != nil {
		fmt.Println("ERROR: TCP server failed to read from client.", read_err)
	}

	reply := svr.requestAction(request)
	_, write_err := conn.Write(reply)
	if write_err != nil {
		fmt.Println("ERROR: TCP server failed to write to client.", write_err)
	}
	conn.Close()
}

func (svr *TCPServer) requestAction(reqBytes []byte) []byte {
	req := raft_model.RpcRequest{}
	if errUnmarshall := raft_util.UnGob(reqBytes, &req); errUnmarshall != nil {
		return errBody("raft:server malformed request", errUnmarshall)
	}
	switch req.Action {
	case raft_model.RequestActions["stats"]:
		return svr.Stats()
	case raft_model.RequestActions["join"]:
		return svr.Join(req.Body)
	case raft_model.RequestActions["leave"]:
		return svr.Leave(req.Body)
	case raft_model.RequestActions["payload"]:
		return svr.Payload(req.Body)
	case raft_model.RequestActions["send-join"]:
		return svr.SendJoin(req.Body)
	case raft_model.RequestActions["send-leave"]:
		return svr.SendLeave(req.Body)
	case raft_model.RequestActions["send-payload"]:
		return svr.SendPayload(req.Body)
	default:
		errAction := fmt.Errorf("unidentified request: %v", req)
		return errBody("raft:server unhandled action", errAction)
	}
}

func errBody(body string, err error) []byte {
	response := raft_model.ServerResponse{
		ErrCode: raft_controller.RaftError["bad-node"],
		Error:   fmt.Errorf("%s: %s", body, err.Error()),
	}
	blob, errGlob := raft_util.Gob(response)
	if errGlob != nil {
		return []byte(errGlob.Error())
	}
	return blob
}
