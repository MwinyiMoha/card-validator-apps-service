package api

import (
	"card-validator-apps-service/internal/core/application"
	"card-validator-apps-service/internal/core/domain"
	protos "card-validator-apps-service/internal/gen"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *Server) GetApps(ctx context.Context, req *empty.Empty) (*protos.GetAppsResponse, error) {
	docs, err := s.Service.ListApps("dummy")
	if err != nil {
		return nil, ParseError(err)
	}

	var apps []*protos.App
	for _, doc := range docs {
		apps = append(apps, MapDocToResponse(doc))
	}

	return &protos.GetAppsResponse{Apps: apps}, nil
}

func (s *Server) GetApp(ctx context.Context, req *protos.GetAppRequest) (*protos.App, error) {
	app, err := s.Service.FetchApp("dummy", req.AppId)
	if err != nil {
		return nil, ParseError(err)
	}

	return MapDocToResponse(app), nil
}

func (s *Server) CreateApp(ctx context.Context, req *protos.CreateAppRequest) (*protos.App, error) {
	payload := &domain.AppPayload{
		OwnerID:     "dummy",
		Name:        req.App.Name,
		Description: req.App.Description,
		Environment: req.App.Environment.String(),
		OwnerType:   req.App.OwnerType.String(),
	}

	app, err := s.Service.CreateApp(payload)
	if err != nil {
		return nil, ParseError(err)
	}

	return MapDocToResponse(app), nil
}

func (s *Server) RefreshAppKey(ctx context.Context, req *protos.RefreshAppKeyRequest) (*protos.RefreshAppKeyResponse, error) {
	newKey, err := s.Service.RefreshKey("dummy", req.AppId)
	if err != nil {
		return nil, ParseError(err)
	}

	return &protos.RefreshAppKeyResponse{NewKey: newKey}, nil
}

func (s *Server) DecodeAppKey(ctx context.Context, req *protos.DecodeAppKeyRequest) (*protos.DecodeAppKeyResponse, error) {
	appId, err := s.Service.DecodeKey(req.AppKey)
	if err != nil {
		return nil, ParseError(err)
	}

	return &protos.DecodeAppKeyResponse{AppId: appId}, nil
}

func (s *Server) DeleteApp(ctx context.Context, req *protos.DeleteAppRequest) (*emptypb.Empty, error) {
	if err := s.Service.DeleteApp("dummy", req.AppId); err != nil {
		return nil, ParseError(err)
	}

	return &emptypb.Empty{}, nil
}
