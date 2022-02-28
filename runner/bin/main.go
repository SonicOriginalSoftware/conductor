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
	if outlog, errlog, listener, grpcServer, err = lib.Init(); err != nil {
		log.Fatalf("%v", err)
	}
}

func main() {
	generated.RegisterRunnerServer(grpcServer, runner.NewService())

	if err = lib.Main(outlog, errlog, listener, grpcServer); err != nil {
		errlog.Fatalf("%v", err)
	}
}
