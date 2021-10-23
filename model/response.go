package raft_model

type ServerResponse struct {
	Message string
	Data    interface{}
	ErrCode int
	Error   error
}
