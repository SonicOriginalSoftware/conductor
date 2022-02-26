package queue

import (
	"conductor/generated"
	"context"
)

// Service contains queue service properties
type Service struct {
	ctx    *context.Context
	cancel *context.CancelFunc

	generated.UnimplementedQueueServer
}

// NewService returns a new Service
//
// ctx serves to scope all service requests to the
// lifetime of the creator of the Service
func NewService(ctx *context.Context, cancel *context.CancelFunc) *Service {
	return &Service{
		ctx:    ctx,
		cancel: cancel,
	}
}
