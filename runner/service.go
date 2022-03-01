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
	info generated.RunnerInfo

	jobCancelToken chan bool
}

func (service *Service) report(result *generated.JobResult) {
	// FIXME Implement reporting back to the queue client

	service.jobCancelToken = nil
	service.CurrentCommandName = ""
	service.CurrentJobName = ""
}

func (service *Service) runCommands(commands []*generated.Command, env []string) (jobResult *generated.JobResult) {
	service.jobCancelToken = make(chan bool, 1)
	defer close(service.jobCancelToken)

	workingDirectory := ""
	jobCancelled := false

	for _, eachCommand := range commands {
		result := &generated.CommandResult{
			Code:    notStarted,
			Message: "Job has not been started",
			// FIXME cmd stderr and stdout need to be appended to files in the working directory
			// stdout and stderr will be appended as each successive command is executed
		}

		service.CurrentCommandName = eachCommand.Name
		cmd := exec.Command(eachCommand.Command)
		cmd.Env = env
		cmd.Dir = workingDirectory

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
		case jobCancelled = <-service.jobCancelToken:
			result.Code = cancelled
			result.Message = "Job was cancelled"
		}

		jobResult.Results = append(jobResult.Results, result)
		if jobCancelled {
			return
		}
	}

	return
}

// Start a Job
func (service *Service) Start(_ context.Context, job *generated.Job) (n *generated.Nil, err error) {
	if service.jobCancelToken != nil {
		return n, fmt.Errorf("Runner already running job: %v", job.Name)
	}

	service.CurrentJobName = job.Name

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
	return &service.info, err
}

// NewService returns a new Service
func NewService() (service *Service, err error) {
	service = &Service{
		info: generated.RunnerInfo{
			Attributes: &generated.Attributes{},
		},
	}

	arch, found := generated.Attributes_Arch_value[runtime.GOARCH]
	if !found {
		return service, fmt.Errorf("Could not obtain Runner Arch")
	}
	service.info.Attributes.Arch = generated.Attributes_Arch(arch)

	platform, found := generated.Attributes_Platform_value[runtime.GOOS]
	if !found {
		return service, fmt.Errorf("Could not obtain Runner Platform")
	}
	service.info.Attributes.Platform = generated.Attributes_Platform(platform)

	switch runtime.GOOS {
	case "windows":
		service.info.Attributes.Libc = generated.Attributes_msvc
	case "darwin":
		service.info.Attributes.Libc = generated.Attributes_libSystem
	case "linux":
		// FIXME Don't assume glibc - check for musl
		service.info.Attributes.Libc = generated.Attributes_glibc
	}

	return
}
