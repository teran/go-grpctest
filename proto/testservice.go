package proto

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
