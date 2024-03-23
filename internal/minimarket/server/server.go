package server

import (
	"context"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	logger *log.Logger
	server *http.Server
}

func NewHttpServer(logLevel, port string, handler http.Handler) (*Server, error) {
	logger, err := configureLogger(logLevel)
	if err != nil {
		return nil, fmt.Errorf("failed to configure logger on httpServer: %w", err)
	}

	apiServer := &Server{
		logger: logger,
		server: &http.Server{
			Addr:    ":" + port,
			Handler: handler,
		},
	}

	return apiServer, nil
}

func (s *Server) Start() error {
	s.logger.Info("Server starting...")
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) {
	err := s.server.Shutdown(ctx)
	if err != nil {
		log.Errorf("error on shutdown server: %s", err.Error())
	}
}

func configureLogger(logLevel string) (*log.Logger, error) {
	logger := log.New()

	lvl, err := log.ParseLevel(logLevel)
	if err != nil {
		return nil, fmt.Errorf("failed to parse level: %w", err)
	}
	logger.SetLevel(lvl)

	return logger, nil
}
