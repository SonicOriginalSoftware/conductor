package main

import (
	"conductor/generated"
	"conductor/lib"
	"conductor/queue"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	outlog        *log.Logger
	errlog        *log.Logger
	listener      net.Listener
	grpcServer    *grpc.Server
	parentContext context.Context
	cancel        context.CancelFunc
	err           error
)

func init() {
	if outlog, errlog, listener, grpcServer, parentContext, cancel, err = lib.Init(); err != nil {
		log.Fatalf("%v", err)
	}
}

func main() {
	generated.RegisterQueueServer(grpcServer, queue.NewService(&parentContext, &cancel))

	if err = lib.Main(parentContext, outlog, errlog, listener, grpcServer, cancel); err != nil {
		errlog.Fatalf("%v", err)
	}
}
