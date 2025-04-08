package crontab

import (
	"context"
	"fmt"
	"time"

	"my-service/model"

	"github.com/appootb/substratum/v2/logger"
)

var (
	gIncr = 0
)

type Broadcast struct{}

func (c Broadcast) Execute(ctx context.Context, _ interface{}) error {
	logger.Info("crontab.Broadcast", logger.Content{
		"timestamp": time.Now(),
	})
	//
	gIncr++
	model.Broadcast(fmt.Sprintf("%s: %d", time.Now().Format("2006-01-02 15:04:05"), gIncr))
	return nil
}
