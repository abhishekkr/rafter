package raft_cfg

import (
	"github.com/abhishekkr/gol/golenv"
)

var (
	ID = golenv.OverrideIfEnv("RAFTER_NODE_ID", "alpha")
	Address = golenv.OverrideIfEnv("RAFTER_NODE_ADDRESS", "127.0.0.1")
	Port = golenv.OverrideIfEnv("RAFTER_NODE_PORT", "6660")
	ServerPort = golenv.OverrideIfEnv("RAFTER_SERVER_PORT", "6661")

	TargetAddress = golenv.OverrideIfEnv("RAFTER_TARGET_ADDRESS", "127.0.0.1")
	TargetPort = golenv.OverrideIfEnv("RAFTER_TARGET_PORT", "6671")
)
