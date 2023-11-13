package grpcslog

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
)

func NewUnaryServerInterceptor(log *slog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		defer func(t time.Time) {
			attrs := []any{
				slog.String("method", info.FullMethod),
				slog.Duration("duration", time.Since(t)),
			}

			if err != nil {
				attrs = append(attrs, slog.String("error", err.Error()))
			}

			log.InfoContext(ctx, "request", attrs...)
		}(time.Now())

		resp, err = handler(ctx, req)
		return
	}
}
