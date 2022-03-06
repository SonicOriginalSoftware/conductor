package lib

import (
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

// Main wraps the main functionality of all servers
func Main(
	outlog *log.Logger,
	errlog *log.Logger,
	listener net.Listener,
	grpcServer *grpc.Server,
	serviceName string,
) (err error) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	served := make(chan error, 1)
	go func() { served <- grpcServer.Serve(listener) }()
	defer close(served)

	outlog.Printf("%v running: %v", serviceName, listener.Addr())

	interrupted := false
	for !interrupted {
		select {
		case err = <-served:
			outlog.Printf("%v stopped: %v\n", serviceName, err)
		case done := <-interrupt:
			outlog.Printf("%v stop requested: %v\n", serviceName, done)
			outlog.Printf("Gracefully shutting down %v...\n", serviceName)
			grpcServer.GracefulStop()
			outlog.Printf("Gracefully shut down %v!\n", serviceName)
		}
		interrupted = true
	}

	outlog.Printf("Service stopped!")
	return
}
