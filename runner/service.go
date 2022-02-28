package runner

import (
	"conductor/generated"
	"context"
	"fmt"
	"os/exec"
	"runtime"
)

type status = int64

const (
	cancelled status = iota
	notStarted
	running
	errored
	awaiting
)

// Service contains Runner service properties
type Service struct {
	generated.UnimplementedRunnerServer
	generated.RunnerStatus

	jobCancelToken chan bool
}

func (service *Service) report(result *generated.JobResult) {
	// FIXME Implement reporting back to the queue client

	service.jobCancelToken = nil
}

func (service *Service) runCommands(commands []*generated.Command, env []string) (result *generated.JobResult) {
	service.jobCancelToken = make(chan bool, 1)
	defer close(service.jobCancelToken)

	result.Code = notStarted
	result.Message = "Job has not been started"

	workingDirectory := ""

	for _, eachCommand := range commands {
		service.CurrentCommandName = eachCommand.Name
		cmd := exec.Command(eachCommand.Command)
		cmd.Env = env
		cmd.Dir = workingDirectory

		// FIXME cmd stderr and stdout need to be appended to files in the working directory
		// stdout and stderr will be appended as each successive command is executed

		runError := make(chan error)
		defer close(runError)
		go func() { runError <- cmd.Run() }()

		select {
		case err := <-runError:
			if exitError, ok := err.(*exec.ExitError); ok {
				result.Code = int64(exitError.ExitCode())
				result.Message = exitError.Error()
			} else {
				result.Code = errored
				result.Message = err.Error()
			}
		case <-service.jobCancelToken:
			result.Code = cancelled
			result.Message = "Job was cancelled"
			return
		}
	}

	return
}

// Start a Job
func (service *Service) Start(_ context.Context, job *generated.Job) (n *generated.Nil, err error) {
	service.CurrentJobName = job.Name

	if service.jobCancelToken != nil {
		return n, fmt.Errorf("Runner already running job: %v", job.Name)
	}

	go func() {
		result := service.runCommands(job.Commands, job.Env)
		service.report(result)
	}()
	return
}

// Stop the Runner's Job
func (service *Service) Stop(context.Context, *generated.Nil) (n *generated.Nil, err error) {
	if service.jobCancelToken == nil {
		return n, fmt.Errorf("Runner is not currently running a job")
	}

	service.jobCancelToken <- true
	return
}

// Status of the Runner's current Job
func (service *Service) Status(context.Context, *generated.Nil) (status *generated.RunnerStatus, err error) {
	return &generated.RunnerStatus{
		CurrentJobName:     service.CurrentJobName,
		CurrentCommandName: service.CurrentCommandName,
	}, err
}

// Info about the Runner
func (service *Service) Info(context.Context, *generated.Nil) (info *generated.RunnerInfo, err error) {
	arch, found := generated.Attributes_Arch_value[runtime.GOARCH]
	if !found {
		return info, fmt.Errorf("Could not obtain Runner Arch")
	}
	info.Attributes.Arch = generated.Attributes_Arch(arch)

	platform, found := generated.Attributes_Platform_value[runtime.GOOS]
	if !found {
		return info, fmt.Errorf("Could not obtain Runner Platform")
	}
	info.Attributes.Platform = generated.Attributes_Platform(platform)

	if runtime.GOOS == "windows" {
		info.Attributes.Libc = generated.Attributes_msvc
	} else {
		// FIXME Don't assume glibc - check for musl
		info.Attributes.Libc = generated.Attributes_glibc
	}

	return
}

// NewService returns a new Service
func NewService() *Service {
	return &Service{}
}
