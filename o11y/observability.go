package o11y

import "log"

func Log(msg interface{}) {
	log.Printf("log: %v", msg)
}

func Metric(vector interface{}) {
	log.Printf("vector: %v", vector)
}
