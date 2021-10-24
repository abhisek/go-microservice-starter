package main

import (
	"context"
	"fmt"
	"net"
	"os"

	api "github.com/abhisek/go-microservice-starter/gen"

	"github.com/abhisek/go-microservice-starter/pkg/domain/pets"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"

	log "github.com/sirupsen/logrus"
)

func buildServer() *grpc.Server {
	server := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_validator.StreamServerInterceptor(),
			),
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_validator.UnaryServerInterceptor(),
			),
		),
	)

	return server
}

func serve(_ *cobra.Command, args []string) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to create tcp listener: %v", err)
	}

	defer listener.Close()

	service, err := pets.NewPetService()
	if err != nil {
		log.Fatalf("Failed to create pet service instance: %v", err)
	}

	server := buildServer()
	api.RegisterPetServiceServer(server, service)

	reflection.Register(server)

	log.Infof("Starting gRPC server on: %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC listener: %v", err)
	}
}

func main() {
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("%s <ConfigFile>", os.Args[0]),
		Short: "Example microservice starter in Go",
		Args:  cobra.ArbitraryArgs,
		Run:   serve,
	}

	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
