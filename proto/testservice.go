package proto

//go:generate protoc -I=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative testservice.proto

func NewTestRequest(s string) *TestRequest {
	return &TestRequest{
		Message: s,
	}
}

func NewTestResponse(s string, err error) (*TestResponse, error) {
	return &TestResponse{
		Message: s,
	}, err
}
