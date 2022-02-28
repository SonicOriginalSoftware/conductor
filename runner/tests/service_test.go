package runner

import (
	"conductor/generated"
	"conductor/runner"
	"context"
	"testing"
)

func TestStartJob(t *testing.T) {
	testJobName := "Test Job"
	testJobCommands := []*generated.Command{}
	testPWDCommand := &generated.Command{
		Name:    "PWD",
		Command: "pwd",
	}
	testWhoAmICommand := &generated.Command{
		Name:    "Who Am I?",
		Command: "whoami",
	}
	testJobCommands = append(testJobCommands, testPWDCommand, testWhoAmICommand)
	testJobEnv := []string{"TEST=true", "FOO=BAR"}

	testJob := &generated.Job{
		Name:     testJobName,
		Commands: testJobCommands,
		Env:      testJobEnv,
	}
	runner := &runner.Service{}

	_, err := runner.Start(context.Background(), testJob)
	if err != nil {
		t.Errorf("%v", err)
	}
}
