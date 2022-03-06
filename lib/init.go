package lib

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

// Init initializes shared variables among all Services
func Init(serviceName string) (
	outlog *log.Logger,
	errlog *log.Logger,
	listener net.Listener,
	grpcServer *grpc.Server,
	err error,
) {
	outlog = log.New(os.Stdout, "", log.LstdFlags)
	errlog = log.New(os.Stderr, "[ERROR] ", log.LstdFlags)

	outlog.Printf("Spinning up %v...", serviceName)

	listener, err = net.Listen("tcp", fmt.Sprintf(":%v", lookupPort()))
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return
	}

	grpcServer = grpc.NewServer([]grpc.ServerOption{}...)

	return
}
