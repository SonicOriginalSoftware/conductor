package queue

import (
	"conductor/generated"
	"context"
)

// Name is the name of the service
const Name = "Queue"

// Service contains queue service properties
type Service struct {
	generated.UnimplementedQueueServer
	info generated.QueueInfo
	Name string

	runners *generated.ConfiguredRunners
}

// Push a pipeline into the queue
func (service *Service) Push(
	ctx context.Context,
	pipeline *generated.Pipeline,
) (_ *generated.Nil, err error) {
	// TODO From the list of configured runners for the queue,
	// triage applicable jobs across them

	return
}

// Finish a Job
func (service *Service) Finish(
	ctx context.Context,
	jobResult *generated.JobResult,
) (_ *generated.Nil, err error) {
	return
}

// Runners are the configured runners reachable by this Queue
func (service *Service) Runners(
	ctx context.Context,
	_ *generated.Nil,
) (runners *generated.ConfiguredRunners, err error) {
	return service.runners, err
}

// NewService returns a new Service
func NewService(address string, runnerAddresses []string) (service *Service, err error) {
	runners := &generated.ConfiguredRunners{}

	for range runnerAddresses {
		// TODO Attempt to communicate with each runner before allowing it to be added
		// e.g. Call the Info or Status method on the RunnerClient
		// And add the results as a ConfiguredRunner
		runners.Runners = append(runners.Runners, &generated.ConfiguredRunner{})
	}

	return &Service{
		Name:    Name,
		runners: runners,
		info: generated.QueueInfo{
			Address: address,
		},
	}, err
}
