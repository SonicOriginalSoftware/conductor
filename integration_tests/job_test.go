package integrationtests

import (
	"conductor/generated"
	"conductor/queue"
	"conductor/runner"
	"context"
	"testing"
)

func setupRunnerService(t *testing.T) (service *runner.Service) {
	service, err := runner.NewService()
	if err != nil {
		t.Errorf("Failed to set up Runner Service:\n%v", err)
	}

	return
}

func setupQueueService(t *testing.T) (service *queue.Service) {
	service, err := queue.NewService()
	if err != nil {
		t.Errorf("Failed to set up Runner Service:\n%v", err)
	}

	return
}

func TestStartJob(t *testing.T) {
	jobName := "Job"
	commandName := "PWD"

	initialJob := &generated.Job{
		Name: jobName,
		Commands: []*generated.Command{
			{
				Name:    commandName,
				Command: "pwd",
			},
		},
	}

	runner := setupRunnerService(t)
	_ = setupQueueService(t)

	if _, err := runner.Start(context.Background(), initialJob); err != nil {
		t.Errorf("%v", err)
	}
}

func TestStartMultipleJobs(t *testing.T) {
	initialJobName := "Initial Job"
	initialCommandName := "Initial Command"
	additionalJobName := "Additional Job"
	additionalCommandName := "Additional Command"

	initialJob := &generated.Job{
		Name: initialJobName,
		Commands: []*generated.Command{
			{
				Name:    initialCommandName,
				Command: "sleep 30",
			},
		},
	}

	additionalJob := &generated.Job{
		Name: additionalJobName,
		Commands: []*generated.Command{
			{
				Name:    additionalCommandName,
				Command: "sleep 30",
			},
		},
	}

	runner := setupRunnerService(t)
	_ = setupQueueService(t)

	if _, err := runner.Start(context.Background(), initialJob); err != nil {
		t.Errorf("%v", err)
	}

	if _, err := runner.Start(context.Background(), additionalJob); err != nil {
		t.Errorf("%v", err)
	}
}
