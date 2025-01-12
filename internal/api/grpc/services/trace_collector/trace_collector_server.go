package trace_collector

import (
	"context"
	"log/slog"
	gen "slogger/internal/api/grpc/gen/services/trace_collector_gen"
)

type Server struct {
	gen.UnimplementedTraceCollectorServer
}

func NewServer() (*Server, error) {
	return &Server{}, nil
}

func (c *Server) Create(ctx context.Context, in *gen.TraceCreateRequest) (*gen.TraceCollectorResponse, error) {
	// TODO: Implement this method
	slog.Info("Creating trace...")

	return &gen.TraceCollectorResponse{StatusCode: 200, Message: "ok"}, nil
}

func (c *Server) Update(ctx context.Context, in *gen.TraceUpdateRequest) (*gen.TraceCollectorResponse, error) {
	// TODO: Implement this method
	slog.Info("Updating trace...")

	return &gen.TraceCollectorResponse{StatusCode: 200, Message: "ok"}, nil
}

func (c *Server) Close() error {
	slog.Warn("grpc: closing trace collector server...")

	return nil
}
