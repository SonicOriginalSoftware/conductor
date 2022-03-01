package runner

import (
	"conductor/generated"
	"context"
	"runtime"
	"testing"
)

func TestRunnerInfo(t *testing.T) {
	runner, err := NewService()
	if err != nil {
		t.Errorf("%v", err)
	}

	info, err := runner.Info(context.Background(), &generated.Nil{})
	if err != nil {
		t.Errorf("%v", err)
	}

	if info.Attributes.Arch != generated.Attributes_Arch(generated.Attributes_Arch_value[runtime.GOARCH]) {
		t.Errorf("Runner Info Arch not consistent with Machine Arch")
		t.FailNow()
	}

	if info.Attributes.Platform != generated.Attributes_Platform(generated.Attributes_Platform_value[runtime.GOOS]) {
		t.Errorf("Runner Info Platform not consistent with Machine OS")
		t.FailNow()
	}

	libc := generated.Attributes_musl
	switch runtime.GOOS {
	case "windows":
		libc = generated.Attributes_msvc
	case "darwin":
		libc = generated.Attributes_libSystem
	case "linux":
		libc = generated.Attributes_glibc
	}
	if info.Attributes.Libc != libc {
		t.Errorf("Runner Info LibC not consistent with Machine LibC")
		t.FailNow()
	}
}

func TestRunnerStatusNoJob(t *testing.T) {
	runner, err := NewService()
	if err != nil {
		t.Errorf("%v", err)
	}

	status, err := runner.Status(context.Background(), &generated.Nil{})
	if err != nil {
		t.Errorf("%v", err)
	}

	if status.CurrentJobName != "" {
		t.Errorf("Runner with no job reported having a current job name")
		t.FailNow()
	}

	if status.CurrentCommandName != "" {
		t.Errorf("Runner with no job reported having a current command name")
		t.FailNow()
	}
}

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

	runner, err := NewService()
	if err != nil {
		t.Errorf("%v", err)
	}

	if _, err := runner.Start(context.Background(), testJob); err != nil {
		t.Errorf("%v", err)
	}
}

func TestStopNoJob(t *testing.T) {
	runner, err := NewService()
	if err != nil {
		t.Errorf("%v", err)
	}

	if _, err := runner.Stop(context.Background(), &generated.Nil{}); err == nil {
		t.Errorf("Runner did not encounter error when trying to stop a job that didn't exist")
		t.FailNow()
	}
}
