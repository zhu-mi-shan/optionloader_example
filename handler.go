package main

import (
	"context"
	example "github.com/zhu-mi-shan/optionloader_example/kitex_gen/example"
)

// TestServiceImpl implements the last service interface defined in the IDL.
type TestServiceImpl struct{}

// Test implements the TestServiceImpl interface.
func (s *TestServiceImpl) Test(ctx context.Context, req *example.Req) (resp *example.Resp, err error) {
	// TODO: Your code here...
	resp = &example.Resp{
		Code: "200",
		Msg:  "ok",
	}

	return
}
