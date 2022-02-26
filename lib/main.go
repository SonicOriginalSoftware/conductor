package lib

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

// Main wraps the main functionality of all servers
func Main(
	parentContext context.Context,
	outlog *log.Logger,
	errlog *log.Logger,
	listener net.Listener,
	grpcServer *grpc.Server,
	cancel context.CancelFunc,
) (err error) {
	defer cancel()

	interrupt, _ := signal.NotifyContext(parentContext, os.Interrupt, os.Kill)

	served := make(chan error, 1)
	go func() { served <- grpcServer.Serve(listener) }()
	defer close(served)

	interrupted := false
	for !interrupted {
		select {
		case err = <-served:
			outlog.Printf("Server stopped: %v\n", err)
		case done := <-interrupt.Done():
			outlog.Printf("Service stop requested: %v\n", done)
			outlog.Printf("Gracefully shutting down server...\n")
			grpcServer.GracefulStop()
			outlog.Printf("Gracefully shut down server!\n")
		}
		interrupted = true
	}

	outlog.Printf("Service stopped!")
	return
}
