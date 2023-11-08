package main

import (
	"log/slog"
	"net"
	"os"

	"github.com/aintsashqa/pathfinder-service/api/proto"
	"github.com/aintsashqa/pathfinder-service/internal/delivery"
	"github.com/aintsashqa/pathfinder-service/internal/handler"
	"google.golang.org/grpc"
)

func main() {
	var (
		hdlr     = handler.New()
		svc      = delivery.NewGrpcService(hdlr)
		srv      = grpc.NewServer()
		lis, err = net.Listen("tcp", ":8080")
	)

	if err != nil {
		slog.Error("unable to create listener", slog.String("error", err.Error()))
		os.Exit(1)
	}

	proto.RegisterPathfinderServiceServer(srv, svc)

	slog.Info("starting server", slog.String("addr", lis.Addr().String()))
	if err := srv.Serve(lis); err != nil {
		slog.Error("unable to serve", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
