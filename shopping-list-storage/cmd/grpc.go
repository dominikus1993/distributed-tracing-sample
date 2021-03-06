package cmd

import (
	"context"
	"flag"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"

	"github.com/dominikus1993/distributed-tracing-sample/shopping-list-storage/internal/handlers"
	"github.com/dominikus1993/distributed-tracing-sample/shopping-list-storage/shoppinglist"
	"github.com/google/subcommands"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
)

type RunGrpcServer struct {
}

func (*RunGrpcServer) Name() string     { return "run-grpc" }
func (*RunGrpcServer) Synopsis() string { return "run grpc server" }
func (*RunGrpcServer) Usage() string {
	return `go run . run-grpc`
}

func (p *RunGrpcServer) SetFlags(f *flag.FlagSet) {
}

func (p *RunGrpcServer) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return subcommands.ExitFailure
	}
	server, err := handlers.InitRpc(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer server.Close(ctx)
	si := grpctrace.StreamServerInterceptor(grpctrace.WithServiceName("shopping-list-storage-api"))
	ui := grpctrace.UnaryServerInterceptor(grpctrace.WithServiceName("shopping-list-storage-api"))
	grpcServer := grpc.NewServer(grpc.StreamInterceptor(si), grpc.UnaryInterceptor(ui))
	reflection.Register(grpcServer)
	shoppinglist.RegisterShoppingListsStorageServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
