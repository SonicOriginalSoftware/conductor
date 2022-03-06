package integrationtests

import (
	"conductor/generated"
	"conductor/lib"
	"conductor/queue"
	"conductor/runner"
	"context"
	"os"
	"testing"
)

const queuePort = "8080"
const runnerPort = "8081"

func setupQueueService(t *testing.T) (service *queue.Service, mainErr chan error) {
	os.Setenv("PORT", queuePort)
	checkPort, found := os.LookupEnv("PORT")
	if !found || checkPort != queuePort {
		t.Fatalf("Queue port could not be set!")
	}

	outlog, errlog, listener, grpcServer, err := lib.Init(queue.Name)
	if err != nil {
		t.Fatalf("%v", err)
	}

	service, err = queue.NewService(listener.Addr().String(), []string{})
	if err != nil {
		t.Errorf("Failed to set up Queue Service:\n%v", err)
	}

	generated.RegisterQueueServer(grpcServer, service)

	mainErr = make(chan error, 1)
	go func() {
		if err = lib.Main(outlog, errlog, listener, grpcServer, service.Name); err != nil {
			mainErr <- err
		}
	}()
	return
}

func setupRunnerService(t *testing.T) (service *runner.Service, mainErr chan error) {
	os.Setenv("PORT", runnerPort)
	checkPort, found := os.LookupEnv("PORT")
	if !found || checkPort != runnerPort {
		t.Fatalf("Runner port could not be set!")
	}

	outlog, errlog, listener, grpcServer, err := lib.Init(runner.Name)
	if err != nil {
		t.Fatalf("%v", err)
	}

	service, err = runner.NewService(listener.Addr().String())
	if err != nil {
		t.Errorf("Failed to set up Runner Service:\n%v", err)
	}

	generated.RegisterRunnerServer(grpcServer, service)

	mainErr = make(chan error, 1)
	go func() {
		if err = lib.Main(outlog, errlog, listener, grpcServer, service.Name); err != nil {
			mainErr <- err
		}
	}()
	return
}

func TestPipelinePushJob(t *testing.T) {
	jobName := "Job"
	commandName := "PWD"

	testJob := &generated.Job{
		Name: jobName,
		Commands: []*generated.Command{
			{
				Name:    commandName,
				Command: "pwd",
			},
		},
	}

	testPipeline := &generated.Pipeline{
		Jobs: []*generated.Job{testJob},
	}

	queue, queueErr := setupQueueService(t)
	defer close(queueErr)

	runner, runnerErr := setupRunnerService(t)
	defer close(runnerErr)

	queue.Push(context.Background(), testPipeline)
	t.Logf("Waiting on %v to finish the job...", runner.Name)

	runnerStatus, err := runner.Status(context.Background(), &generated.Nil{})
	if err != nil {
		t.Fatalf("Could not retrieve runner status")
	}

	if runnerStatus.CurrentJobName == "" {
		t.Fatalf("Runner reported as not having a job")
	}

	// FIXME What am I waiting on here?
	select {
	case err := <-runnerErr:
		t.Fatalf("%v", err)
		break
	case err := <-queueErr:
		t.Fatalf("%v", err)
		break
	case _ = <-runner.WaitForJob():
		break
	}

	t.Logf("Able to push pipeline with single job!")
}

func TestPipelinePushJobs(t *testing.T) {
	t.Skip("Not ready to test multiple jobs yet")

	initialJobName := "Initial Job"
	initialCommandName := "Initial Command"
	additionalJobName := "Additional Job"
	additionalCommandName := "Additional Command"

	_ = &generated.Job{
		Name: initialJobName,
		Commands: []*generated.Command{
			{
				Name:    initialCommandName,
				Command: "sleep 30",
			},
		},
	}

	_ = &generated.Job{
		Name: additionalJobName,
		Commands: []*generated.Command{
			{
				Name:    additionalCommandName,
				Command: "sleep 30",
			},
		},
	}

	_, _ = setupRunnerService(t)
	_, _ = setupQueueService(t)

	t.Logf("Able to push pipeline with multiple job!")
}
