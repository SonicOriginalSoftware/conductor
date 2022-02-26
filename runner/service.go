package runner

import (
	"conductor/generated"
	"context"
	"os/exec"
)

// Service contains runner service properties
type Service struct {
	ctx    *context.Context
	cancel *context.CancelFunc

	generated.UnimplementedRunnerServer
}

// Run a Job
func (rs *Service) Run(_ context.Context, job *generated.Job) (*generated.Nil, error) {
	workingDirectory := ""

	cmd := exec.Command(job.Command)
	cmd.Env = job.Env
	cmd.Dir = workingDirectory

	// cmd.Stdout =
	// cmd.Stderr =

	// FIXME execute cmd.Run() in a separate goroutine that creates
	// a queue client and calls the finish method on it with the
	// results of the cmd.Run() call
	return &generated.Nil{}, cmd.Run()
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
