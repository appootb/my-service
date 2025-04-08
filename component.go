package my_service

import (
	"context"

	"my-service/config"
	"my-service/crontab"
	example "my-service/protobuf/go"
	"my-service/rpc"

	"github.com/appootb/substratum/v2"
	"github.com/appootb/substratum/v2/configure"
	"github.com/appootb/substratum/v2/queue"
	"github.com/appootb/substratum/v2/service"
	"github.com/appootb/substratum/v2/storage"
	"github.com/appootb/substratum/v2/task"
)

type MyComponent struct {
	context.Context
}

func New(ctx context.Context) substratum.Component {
	return &MyComponent{
		Context: ctx,
	}
}

func (m MyComponent) Name() string {
	return "my_service" // A unique service name for service discovery
}

func (m MyComponent) Init(cfg configure.Configure) error {
	return cfg.Register(m.Name(),
		config.Settings(),
		configure.WithAutoCreation(true))
}

func (m MyComponent) InitStorage(s storage.Storage) error {
	// // Init Redis
	// if err := s.InitRedis(config.Settings().RedisAddresses); err != nil {
	// 	return err
	// }
	// // Init Queue
	// if err := s.InitCommon(config.Settings().QueueAddress); err != nil {
	// 	return err
	// }
	return nil
}

func (m MyComponent) RegisterHandler(outer, inner service.HttpHandler) error {
	return nil
}

func (m MyComponent) RegisterService(auth service.Authenticator, srv service.Implementor) error {
	return example.RegisterMyServiceScopeServer(m.Name(), auth, srv, &rpc.Example{})
}

func (m MyComponent) RunQueueWorker(q queue.Queue) error {
	return nil
}

func (m MyComponent) ScheduleCronTask(t task.Task) error {
	if err := t.Schedule("@every 5s", &crontab.Broadcast{},
		task.WithComponent(m.Name()),
		task.WithSingleton()); err != nil {
		return err
	}
	return nil
}
