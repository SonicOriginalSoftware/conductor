package server

import (
	"conductor/generated"
	"context"
)

// Service contains queue service properties
type Service struct {
	generated.UnimplementedServerServer
}

// NewService returns a new Service
//
// ctx serves to scope all service requests to the
// lifetime of the creator of the Service
func NewService() *Service {
	return &Service{}
}

// Finish a Pipeline
func (service *Service) Finish(
	context.Context,
	*generated.PipelineResult,
) (_ *generated.Nil, err error) {
	return
}
