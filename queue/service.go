package queue

import (
	"conductor/generated"
	"context"
	"log"

	"google.golang.org/grpc"
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
	var availableRunners generated.ConfiguredRunners

	var runnerOptions []grpc.DialOption

	for _, eachRunner := range service.runners.Runners {
		connection, err := grpc.Dial(eachRunner.Info.Address, runnerOptions)
		if err != nil {
			log.Printf("%v", err)
			continue
		}

		defer connection.Close()
		runnerClient := generated.NewRunnerClient(connection)
		runnerStatus, err := runnerClient.Status(ctx, &generated.Nil{})
		if err != nil {
			log.Printf("%v", err)
			continue
		}

		runnerInfo, err := runnerClient.Info(ctx, &generated.Nil{})
		if err != nil {
			log.Printf("%v", err)
			continue
		}

		// TODO Verify Runner Attributes

		// TODO Verify Runner Status
		// if runnerInfo.Attributes.Arch ==

		// FIXME If the runner attributes fit the job
		// AND the runner is listed as available...
		// Then add the runner to the list of available runners
		if true {
			availableRunners.Runners = append(availableRunners.Runners, eachRunner)
		}

	}

	// TODO From the list of configured runners for the queue,
	// determine which runners are applicable and available for jobs,
	// then triage applicable jobs across them

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
