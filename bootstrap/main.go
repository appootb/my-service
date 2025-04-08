package main

import (
	"fmt"
	"log"

	comp "my-service"

	"github.com/appootb/substratum/v2"
	"github.com/appootb/substratum/v2/context"
	"github.com/appootb/substratum/v2/metadata"
	"github.com/appootb/substratum/v2/proto/go/permission"
)

const (
	ClientRpcPort     = 6007
	ClientGatewayPort = 6009
	ServerRpcPort     = 8007
	ServerGatewayPort = 8009
)

func main() {
	// New server instance
	srv := substratum.NewServer(
		substratum.WithServeMux(permission.VisibleScope_CLIENT, ClientRpcPort, ClientGatewayPort),
		substratum.WithServeMux(permission.VisibleScope_SERVER, ServerRpcPort, ServerGatewayPort))

	// Register components
	if err := srv.Register(comp.New(context.Context())); err != nil {
		log.Panicf("register component failed, err: %v", err)
	}

	// Serve
	if err := srv.Serve(metadata.EnvDevelop == "local"); err != nil {
		fmt.Println("exiting...", err.Error())
	}
}
