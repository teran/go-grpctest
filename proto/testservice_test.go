package proto

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestNewTestRequest(t *testing.T) {
	r := require.New(t)

	req := NewTestRequest("blah")
	r.Equal(&TestRequest{
		Message: "blah",
	}, req)
}

func TestNewTestResponse(t *testing.T) {
	r := require.New(t)

	resp, err := NewTestResponse("blah", errors.New("blah"))
	r.Equal("blah", err.Error())
	r.Equal(&TestResponse{
		Message: "blah",
	}, resp)
}
