package integrationtests

import (
	"conductor/generated"
	"conductor/queue"
	"conductor/runner"
	"context"
	"testing"
)

func TestStartJob(t *testing.T) {
	initialJobName := "Initial Job"
	initialCommandName := "Initial Command"

	initialJob := &generated.Job{
		Name: initialJobName,
		Commands: []*generated.Command{
			{
				Name:    initialCommandName,
				Command: "sleep 30",
			},
		},
	}

	runner, err := runner.NewService()
	if err != nil {
		t.Errorf("%v", err)
	}

	_, err = queue.NewService()
	if err != nil {
		t.Errorf("%v", err)
	}

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

	runner, err := runner.NewService()
	if err != nil {
		t.Errorf("%v", err)
	}

	_, err = queue.NewService()
	if err != nil {
		t.Errorf("%v", err)
	}

	if _, err := runner.Start(context.Background(), initialJob); err != nil {
		t.Errorf("%v", err)
	}

	if _, err := runner.Start(context.Background(), additionalJob); err != nil {
		t.Errorf("%v", err)
	}
}
