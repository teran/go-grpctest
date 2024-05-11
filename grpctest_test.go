package grpctest

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/teran/go-grpctest/proto"
)

func (s *testSuite) TestGRPCTestHelper() {
	s.handlers.
		On(
			"TestMessage",
			"somerequestvalue",
		).
		Return(
			"someresponsevalue",
			nil,
		)

	resp, err := s.client.TestMessage(s.ctx, proto.NewTestRequest("somerequestvalue"))
	s.Require().NoError(err)
	s.Require().Equal("someresponsevalue", resp.GetMessage())
}

// ========================================================================
// Test suite setup
// ========================================================================
type testSuite struct {
	suite.Suite

	handlers *handlersMock
	client   proto.TestServiceClient

	srv Server

	ctx    context.Context
	cancel context.CancelFunc
}

func (s *testSuite) SetupTest() {
	s.handlers = &handlersMock{}
	s.srv = New()
	proto.RegisterTestServiceServer(s.srv.Server(), s.handlers)

	err := s.srv.Run()
	s.Require().NoError(err)

	s.ctx, s.cancel = context.WithTimeout(context.Background(), 10*time.Second)

	dial, err := s.srv.DialContext(s.ctx)
	s.Require().NoError(err)

	s.client = proto.NewTestServiceClient(dial)
}

func (s *testSuite) TearDownTest() {
	s.srv.Close()
	s.cancel()

	s.handlers.AssertExpectations(s.T())
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, &testSuite{})
}
