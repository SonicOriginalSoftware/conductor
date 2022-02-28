package lib

import (
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

// Main wraps the main functionality of all servers
func Main(outlog *log.Logger, errlog *log.Logger, listener net.Listener, grpcServer *grpc.Server) (err error) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	served := make(chan error, 1)
	go func() { served <- grpcServer.Serve(listener) }()
	defer close(served)

	outlog.Printf("Service running: %v", listener.Addr())

	interrupted := false
	for !interrupted {
		select {
		case err = <-served:
			outlog.Printf("Server stopped: %v\n", err)
		case done := <-interrupt:
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
