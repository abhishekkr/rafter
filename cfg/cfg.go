package raft_cfg

import (
	"time"

	"github.com/abhishekkr/gol/golenv"
)

var (
	ID         = golenv.OverrideIfEnv("RAFTER_NODE_ID", "alpha")
	Address    = golenv.OverrideIfEnv("RAFTER_NODE_ADDRESS", "127.0.0.1")
	Port       = golenv.OverrideIfEnv("RAFTER_NODE_PORT", "6660")
	ServerPort = golenv.OverrideIfEnv("RAFTER_SERVER_PORT", "6661")

	TargetAddress = golenv.OverrideIfEnv("RAFTER_TARGET_ADDRESS", "127.0.0.1")
	TargetPort    = golenv.OverrideIfEnv("RAFTER_TARGET_PORT", "6671")

	RaftFsmEngine = golenv.OverrideIfEnv("RAFTER_FSM_ENGINE", "badger")
	RaftVolumeDir = golenv.OverrideIfEnv("RAFTER_VOLUME_DIR", "/tmp/raft-badger")

	// The maxPool controls how many connections we will pool.
	RaftMaxPool = 3

	// The timeout is used to apply I/O deadlines. For InstallSnapshot, we multiply
	// the timeout by (SnapshotSize / TimeoutScale).
	// https://github.com/hashicorp/raft/blob/v1.1.2/net_transport.go#L177-L181
	RaftTcpTimeout = 10 * time.Second

	// The `retain` parameter controls how many
	// snapshots are retained. Must be at least 1.
	RaftSnapShotRetain = 2

	// raftLogCacheSize is the maximum number of logs to cache in-memory.
	// This is used to reduce disk I/O for the recently committed entries.
	RaftLogCacheSize = 512

	SuccessCode = "200"
)
