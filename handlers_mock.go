package grpctest

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/teran/go-grpctest/proto"
)

type handlersMock struct {
	mock.Mock
	proto.UnimplementedTestServiceServer
}

func (m *handlersMock) TestMessage(_ context.Context, tr *proto.TestRequest) (*proto.TestResponse, error) {
	args := m.Called(tr.GetMessage())
	return proto.NewTestResponse(args.Get(0).(string), args.Error(1))
}
