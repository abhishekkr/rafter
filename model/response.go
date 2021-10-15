package raft_model

import (
	"bytes"
	"encoding/gob"
)

type ServerResponse struct {
	Body []byte
	Err  error
}

func (response *ServerResponse) UnGob(responseBody []byte) error {
	buf := bytes.NewBuffer(responseBody)
	return gob.NewDecoder(buf).Decode(response)
}
