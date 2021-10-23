package raft_model

import (
	"github.com/hashicorp/raft"
)

// CommandPayload is payload sent by system when calling raft.Apply(cmd []byte, timeout time.Duration)
type CommandPayload struct {
	Operation string
	Key       string
	Value     interface{}
}

// ApplyResponse response from Apply raft
type ApplyResponse struct {
	Error error
	Data  interface{}
}

// snapshotNoop handle noop snapshot
type SnapshotNoop struct{}

// Persist persist to disk. Return nil on success, otherwise return error.
func (s SnapshotNoop) Persist(_ raft.SnapshotSink) error { return nil }

// Release release the lock after persist snapshot.
// Release is invoked when we are finished with the snapshot.
func (s SnapshotNoop) Release() {}

// NewSnapshotNoop is returned by an FSM in response to a SnapshotNoop
// It must be safe to invoke FSMSnapshot methods with concurrent
// calls to Apply.
func NewSnapshotNoop() (raft.FSMSnapshot, error) {
	return &SnapshotNoop{}, nil
}
