package server

import (
	"conductor/generated"
	"context"
)

// Name is the name of the service
const Name = "Server"

// Service contains queue service properties
type Service struct {
	generated.UnimplementedServerServer

	Name string
}

// NewService returns a new Service
//
// ctx serves to scope all service requests to the
// lifetime of the creator of the Service
func NewService() *Service {
	return &Service{
		Name: Name,
	}
}

// Finish a Pipeline
func (service *Service) Finish(
	context.Context,
	*generated.PipelineResult,
) (_ *generated.Nil, err error) {
	return
}
