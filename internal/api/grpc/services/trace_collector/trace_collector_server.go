package trace_collector

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log/slog"
	gen "slogger/internal/api/grpc/gen/services/trace_collector_gen"
	"strings"
)

const bearerPrefix = "Bearer "

type Server struct {
	gen.UnimplementedTraceCollectorServer
}

func NewServer() (*Server, error) {
	return &Server{}, nil
}

func (c *Server) Create(ctx context.Context, in *gen.TraceCreateRequest) (*gen.TraceCollectorResponse, error) {
	// TODO: Implement this method

	go func() {
		token, err := c.detectBearerToken(ctx)

		if err != nil {
			slog.Error("Error detecting bearer token: " + err.Error())

			return
		}

		slog.Info("Creating trace: " + token)
	}()

	return &gen.TraceCollectorResponse{StatusCode: 200, Message: "ok"}, nil
}

func (c *Server) Update(ctx context.Context, in *gen.TraceUpdateRequest) (*gen.TraceCollectorResponse, error) {
	// TODO: Implement this method

	go func() {
		token, err := c.detectBearerToken(ctx)

		if err != nil {
			slog.Error("Error detecting bearer token: " + err.Error())

			return
		}

		slog.Info("Updating trace: " + token)
	}()

	return &gen.TraceCollectorResponse{StatusCode: 200, Message: "ok"}, nil
}

func (c *Server) Close() error {
	slog.Warn("grpc: closing trace collector server...")

	return nil
}

func (c *Server) detectBearerToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return "", fmt.Errorf("missing metadata")
	}

	authValues := md.Get("authorization")

	if len(authValues) == 0 {
		return "", errors.New("missing authorization")
	}

	authorization := authValues[0]

	if authorization == "" {
		return "", errors.New("missing authorization value")
	}

	if !strings.HasPrefix(authorization, bearerPrefix) {
		return "", errors.New("invalid authorization value")
	}

	token := strings.Trim(strings.TrimPrefix(authorization, bearerPrefix), " ")

	if token == "" {
		return "", errors.New("invalid authorization value")
	}

	return token, nil
}
