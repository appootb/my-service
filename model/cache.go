package model

import (
	"sync"

	example "my-service/protobuf/go"
)

var (
	globalSessions = sync.Map{}
)

func Broadcast(msg string) {
	globalSessions.Range(func(key, value interface{}) bool {
		session := value.(*Session)
		session.PostEvent(&example.DownStream{
			Message: msg,
		})
		return true
	})
}
