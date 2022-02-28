package runner

import (
	"conductor/generated"
	"conductor/runner"
	"context"
	"testing"
)

func TestStartJob(t *testing.T) {
	testJobName := "Test Job"
	testJobCommands := []*generated.Command{
		{
			Name:    "PWD",
			Command: "pwd",
		},
		{
			Name:    "Who Am I?",
			Command: "whoami",
		},
	}
	testJob := &generated.Job{
		Name:     testJobName,
		Commands: testJobCommands,
		Env:      []string{"TEST=true", "FOO=BAR"},
	}
	runner := &runner.Service{}

	if _, err := runner.Start(context.Background(), testJob); err != nil {
		t.Errorf("%v", err)
	}

	t.Logf("Finished %v", testJobName)
}
