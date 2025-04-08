package config

import (
	"github.com/appootb/substratum/v2/configure"
)

var (
	cfg setting
)

type setting struct {
	IntVal     int64         `comment:"Common type" default:"1"`
	DynamicVal configure.Int `comment:"Dynamic type" default:"2"`

	// // Queue address
	// QueueAddress configure.Address `comment:"Queue address" default:""`
	// // Redis
	// RedisAddresses []configure.Address `comment:"Redis addresses" default:"redis://:@127.0.0.1:6379/0"`
}

func Settings() *setting {
	return &cfg
}
