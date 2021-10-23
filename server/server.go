package raft_server

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"

	raft_cfg "github.com/abhishekkr/rafter/cfg"
	raft_controller "github.com/abhishekkr/rafter/controller"
	raft_fsm "github.com/abhishekkr/rafter/fsm"
)

type Manager struct {
	Server *TCPServer
	Store  raft_fsm.StoreHandle
}

func New() *Manager {
	var raftBinAddr = fmt.Sprintf("%s:%s", raft_cfg.Address, raft_cfg.Port)

	raftConf := raft.DefaultConfig()
	raftConf.LocalID = raft.ServerID(raft_cfg.ID)
	raftConf.SnapshotThreshold = 1024

	fsmStore, errFSM := raft_fsm.NewFSM()
	if errFSM != nil {
		log.Fatal(errFSM)
		return &Manager{}
	}

	store, err := raftboltdb.NewBoltStore(filepath.Join(raft_cfg.RaftVolumeDir, "raft.dataRepo"))
	if err != nil {
		log.Fatal(err)
		return &Manager{}
	}

	// Wrap the store in a LogCache to improve performance.
	cacheStore, err := raft.NewLogCache(raft_cfg.RaftLogCacheSize, store)
	if err != nil {
		log.Fatal(err)
		return &Manager{}
	}

	snapshotStore, err := raft.NewFileSnapshotStore(raft_cfg.RaftVolumeDir, raft_cfg.RaftSnapShotRetain, os.Stdout)
	if err != nil {
		log.Fatal(err)
		return &Manager{}
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", raftBinAddr)
	if err != nil {
		log.Fatal(err)
		return &Manager{}
	}

	transport, err := raft.NewTCPTransport(raftBinAddr, tcpAddr, raft_cfg.RaftMaxPool, raft_cfg.RaftTcpTimeout, os.Stdout)
	if err != nil {
		log.Fatal(err)
		return &Manager{}
	}

	raftServer, err := raft.NewRaft(raftConf, fsmStore.FSM, cacheStore, store, snapshotStore, transport)
	if err != nil {
		log.Fatal(err)
		return &Manager{}
	}

	// always start single server as a leader
	configuration := raft.Configuration{
		Servers: []raft.Server{
			{
				ID:      raft.ServerID(raft_cfg.ID),
				Address: transport.LocalAddr(),
			},
		},
	}

	raftServer.BootstrapCluster(configuration)

	thisNode := raft_controller.New(raftServer)
	tcpServer := &TCPServer{nodeHandler: thisNode}
	return &Manager{Server: tcpServer, Store: fsmStore.Store}
}
