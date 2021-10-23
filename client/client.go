package raft_client

func New(connectionString string) TCPClient {
	// prepare client for TargetAddress:TargetPort
	return createTCPClient(connectionString)
}
