package lib

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

var port = lookupPort()

// Init initializes shared variables among all Services
func Init() (
	outlog *log.Logger,
	errlog *log.Logger,
	listener net.Listener,
	grpcServer *grpc.Server,
	err error,
) {
	outlog = log.New(os.Stdout, "", log.LstdFlags)
	errlog = log.New(os.Stderr, "[ERROR] ", log.LstdFlags)

	outlog.Printf("Spinning up service...")

	listener, err = net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return
	}

	grpcServer = grpc.NewServer([]grpc.ServerOption{}...)

	return
}
