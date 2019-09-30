package server

import (
	"context"

	api "bitbucket.org/agrostar/grpc_test/proto"
)

type SomeStruct struct {
	A string
}

func (s *SomeStruct) SomeFunction(ctx context.Context, a *api.MSG1) (*api.MSG1, error) {

	return a, nil
}
