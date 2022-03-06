package main

import (
	"conductor/generated"
	"conductor/lib"
	"conductor/runner"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	outlog     *log.Logger
	errlog     *log.Logger
	listener   net.Listener
	grpcServer *grpc.Server
	err        error
)

func init() {
	if outlog, errlog, listener, grpcServer, err = lib.Init(runner.Name); err != nil {
		log.Fatalf("%v", err)
	}
}

func main() {
	service, err := runner.NewService(listener.Addr().String())
	if err != nil {
		errlog.Fatalf("%v", err)
		return
	}

	generated.RegisterRunnerServer(grpcServer, service)

	if err = lib.Main(outlog, errlog, listener, grpcServer, service.Name); err != nil {
		errlog.Fatalf("%v", err)
	}
}
