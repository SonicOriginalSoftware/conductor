package main

import (
	"conductor/generated"
	"conductor/lib"
	"conductor/server"
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
	if outlog, errlog, listener, grpcServer, err = lib.Init(server.Name); err != nil {
		log.Fatalf("%v", err)
	}
}

func main() {
	service := server.NewService()

	generated.RegisterServerServer(grpcServer, service)

	if err = lib.Main(outlog, errlog, listener, grpcServer, service.Name); err != nil {
		errlog.Fatalf("%v", err)
	}
}
