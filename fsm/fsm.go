package raft_fsm

import (
	"log"

	"github.com/dgraph-io/badger/v2"
	"github.com/hashicorp/raft"

	raft_cfg "github.com/abhishekkr/rafter/cfg"
)

type StoreHandle interface {
	Close() error
}
type Machine struct {
	FSM   raft.FSM
	Store StoreHandle
}

// NewBadger raft.FSM implementation using badgerDB
func NewFSM() (*Machine, error) {
	// Preparing badgerDB
	badgerOpt := badger.DefaultOptions(raft_cfg.RaftVolumeDir)
	badgerDB, err := badger.Open(badgerOpt)
	if err != nil {
		log.Fatal(err)
		return &Machine{&badgerFSM{}, badgerDB}, err
	}
	return &Machine{NewBadger(badgerDB), badgerDB}, nil
}
