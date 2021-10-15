package raft_util

import (
	"bytes"
	"encoding/gob"
)

func Gob(skeleton interface{}) (blob []byte, err error) {
	buf := new(bytes.Buffer)
	if err = gob.NewEncoder(buf).Encode(skeleton); err != nil {
		return
	}
	blob = buf.Bytes()
	return
}

func UnGob(body []byte, skeleton interface{}) error {
	buf := bytes.NewBuffer(body)
	return gob.NewDecoder(buf).Decode(skeleton)
}
