package server

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/deeeerain/nebula-stub/core/logger"
	"github.com/deeeerain/nebula-stub/core/middlewares"
)

type Server struct {
	logger      logger.Logger
	middlewares []middlewares.Middleware
	mux         *http.ServeMux
}

func New(logger logger.Logger) *Server {
	return &Server{
		logger: logger,
		mux:    http.NewServeMux(),
	}
}

func (s *Server) Logger() logger.Logger {
	if s.logger == nil {
		return slog.Default()
	}

	return s.logger
}

func (s *Server) Use(middlewares ...middlewares.Middleware) {
	s.middlewares = append(s.middlewares, middlewares...)
}

func (s *Server) Handle(pattern string, handler http.Handler) {
	s.mux.Handle(pattern, handler)
}

func (s *Server) HandleFunc(pattern string, handler http.HandlerFunc) {
	s.mux.HandleFunc(pattern, handler)
}

func (s *Server) Run(listen string) error {
	listenr, err := net.Listen("tcp", listen)
	if err != nil {
		return fmt.Errorf("Failed to create listener: %w", err)
	}

	s.Logger().Info("Server starting", "bind", listenr.Addr())

	chain := middlewares.Chain(s.mux, s.middlewares...)
	return http.Serve(listenr, chain)
}
