package raft_util

import (
	jsoniter "github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Gob(skeleton interface{}) ([]byte, error) {
	return json.Marshal(skeleton)
}

func UnGob(body []byte, skeleton interface{}) error {
	return json.Unmarshal(body, skeleton)
}
