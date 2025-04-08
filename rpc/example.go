package rpc

import (
	"context"
	"math/rand"
	"time"

	"my-service/model"
	example "my-service/protobuf/go"

	"github.com/appootb/substratum/v2/proto/go/permission"
	"github.com/appootb/substratum/v2/proto/go/secret"
	"github.com/appootb/substratum/v2/token"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Example struct {
	example.UnimplementedMyServiceServer
}

func (s *Example) Login(ctx context.Context, _ *emptypb.Empty) (*example.Token, error) {
	now := time.Now()
	secretInfo := &secret.Info{
		Type:      secret.Type_CLIENT,
		Algorithm: secret.Algorithm_HMAC,
		Issuer:    "my-service",
		Account:   rand.Uint64(),
		Subject:   permission.Subject_WEB,
		IssuedAt:  timestamppb.New(now),
		ExpiredAt: timestamppb.New(now.Add(time.Hour)),
	}
	val, err := token.Implementor().Generate(secretInfo)
	if err != nil {
		return nil, err
	}
	return &example.Token{
		Token: val,
	}, nil
}

func (s *Example) Stream(stream example.MyService_StreamServer) error {
	session := model.NewSession(stream)
	//
	<-session.Context().Done()
	session.Close()
	return nil
}
