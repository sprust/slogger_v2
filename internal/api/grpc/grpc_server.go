package grpc

import (
	"google.golang.org/grpc"
	"io"
	"log/slog"
	"net"
	"slogger/internal/api/grpc/gen/services/trace_collector_gen"
	"slogger/internal/api/grpc/services/trace_collector"
	"slogger/pkg/foundation/errs"
)

type Server struct {
	port       string
	grpcServer *grpc.Server
	servers    []io.Closer
}

func NewServer(rpcPort string) *Server {
	server := &Server{
		port: rpcPort,
	}

	return server
}

func (s *Server) Run() error {
	s.grpcServer = grpc.NewServer()

	err := s.registerTraceCollectorServer(s.grpcServer)

	if err != nil {
		return errs.Err(err)
	}

	listener, err := net.Listen("tcp", ":"+s.port)

	if err != nil {
		slog.Error("grpc: error listening: ", err.Error())

		return errs.Err(err)
	}

	slog.Info("grpc: listening on " + s.port)

	err = s.grpcServer.Serve(listener)

	if err != nil {
		slog.Error("grpc: error serving: ", err.Error())
	}

	return errs.Err(err)
}

func (s *Server) Close() error {
	slog.Warn("grpc: closing grpc server...")

	if s.grpcServer == nil {
		return nil
	}

	for _, server := range s.servers {
		err := server.Close()

		if err != nil {
			return errs.Err(err)
		}
	}

	s.grpcServer.Stop()

	return nil
}

func (s *Server) registerTraceCollectorServer(grpcServer *grpc.Server) error {
	server, err := trace_collector.NewServer()

	if err != nil {
		slog.Error("grpc: error creating collector: ", err.Error())

		return errs.Err(err)
	}

	trace_collector_gen.RegisterTraceCollectorServer(grpcServer, server)

	s.servers = append(s.servers, server)

	return nil
}
