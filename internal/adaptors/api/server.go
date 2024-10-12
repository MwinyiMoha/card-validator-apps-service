package api

import (
	"card-validator-apps-service/internal/core/application"
	protos "card-validator-apps-service/internal/gen"
)

type Server struct {
	protos.UnimplementedAppsServiceServer
	Service application.AppService
}

func NewServer(svc application.AppService) *Server {
	return &Server{
		Service: svc,
	}
}
