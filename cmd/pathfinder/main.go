package main

import (
	"flag"
	"log/slog"
	"net"
	"os"

	"github.com/aintsashqa/pathfinder-service/api/proto"
	"github.com/aintsashqa/pathfinder-service/internal/delivery"
	"github.com/aintsashqa/pathfinder-service/internal/handler"
	grpcslog "github.com/aintsashqa/pathfinder-service/pkg/grpc-slog"
	"google.golang.org/grpc"
)

var (
	dev = flag.Bool("dev", false, "development mode")
)

func init() {
	flag.Parse()
}

func main() {
	logOptions := slog.HandlerOptions{}
	log := slog.New(slog.NewJSONHandler(os.Stdout, &logOptions))
	if *dev {
		log = slog.New(slog.NewTextHandler(os.Stdout, &logOptions))
	}

	var (
		hdlr = handler.New()
		svc  = delivery.NewGrpcService(hdlr)
		srv  = grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				grpcslog.NewUnaryServerIntercepter(log),
			),
		)
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
