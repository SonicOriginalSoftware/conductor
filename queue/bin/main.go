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
	if outlog, errlog, listener, grpcServer, err = lib.Init(queue.Name); err != nil {
		log.Fatalf("%v", err)
	}
}

func main() {
	// FIXME Should get runner addresses from a config somewhere?
	service, err := queue.NewService(listener.Addr().String(), []string{})
	if err != nil {
		errlog.Fatalf("%v", err)
		return
	}

	generated.RegisterQueueServer(grpcServer, service)

	if err = lib.Main(outlog, errlog, listener, grpcServer, service.Name); err != nil {
		errlog.Fatalf("%v", err)
	}
}
