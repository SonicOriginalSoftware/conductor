package queue

import (
	"conductor/generated"
	"context"
)

// Service contains queue service properties
type Service struct {
	generated.UnimplementedQueueServer
}

// NewService returns a new Service
//
// ctx serves to scope all service requests to the
// lifetime of the creator of the Service
func NewService() *Service {
	return &Service{}
}

// Push a pipeline into the queue
func (service *Service) Push(
	ctx context.Context,
	pipeline *generated.Pipeline,
) (_ *generated.Nil, err error) {
	return
}

// Finish a Job
func (service *Service) Finish(
	ctx context.Context,
	jobResult *generated.JobResult,
) (_ *generated.Nil, err error) {
	return
}
