package main

import (
	"conductor/generated"
	"conductor/lib"
	"conductor/queue"
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
	generated.RegisterQueueServer(grpcServer, queue.NewService())

	if err = lib.Main(outlog, errlog, listener, grpcServer); err != nil {
		errlog.Fatalf("%v", err)
	}
}
