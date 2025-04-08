package model

import (
	"context"

	example "my-service/protobuf/go"

	"github.com/appootb/substratum/v2/logger"
	"github.com/appootb/substratum/v2/model"
	"github.com/appootb/substratum/v2/service"
)

type Session struct {
	model.Base

	ch chan *example.DownStream

	uid    uint64
	stopFn context.CancelFunc
	stream example.MyService_StreamServer
}

func NewSession(stream example.MyService_StreamServer) *Session {
	uid := service.AccountSecretFromContext(stream.Context()).GetAccount()
	ctx, cancel := context.WithCancel(stream.Context())
	session := &Session{
		Base:   model.New(model.WithContext(ctx)),
		ch:     make(chan *example.DownStream, 100),
		stopFn: cancel,
		stream: stream,
	}
	go session.Serve()
	//
	if old, ok := globalSessions.Load(uid); ok {
		if os, valid := old.(*Session); valid {
			os.Close()
		}
		globalSessions.Delete(uid)
	}
	globalSessions.Store(uid, session)
	return session
}

func (m *Session) Close() {
	select {
	case <-m.Context().Done():
	default:
		m.stopFn()
		globalSessions.Delete(m.uid)
	}
}

func (m *Session) Serve() {
	for {
		select {
		case <-m.Context().Done():
			m.Close()
			return

		case evt := <-m.ch:
			err := m.stream.Send(evt)
			if err != nil {
				m.Close()
				m.Logger().Error("Session.Serve", logger.Content{
					"Send err": err.Error(),
				})
				return
			}
		}
	}
}

func (m *Session) PostEvent(downstream *example.DownStream) {
	if downstream == nil {
		return
	}

Enqueue:
	for {
		select {
		case m.ch <- downstream:
			break Enqueue
		default:
			<-m.ch
		}
	}
}
