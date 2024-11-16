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
	userId, err := GetAuthUser(ctx)
	if err != nil {
		return nil, ParseError(err).Err()
	}

	docs, err := s.Service.ListApps(userId)
	if err != nil {
		return nil, ParseError(err).Err()
	}

	var apps []*protos.App
	for _, doc := range docs {
		apps = append(apps, MapDocToResponse(doc))
	}

	return &protos.GetAppsResponse{Apps: apps}, nil
}

func (s *Server) GetApp(ctx context.Context, req *protos.GetAppRequest) (*protos.App, error) {
	userId, err := GetAuthUser(ctx)
	if err != nil {
		return nil, ParseError(err).Err()
	}

	app, err := s.Service.FetchApp(userId, req.AppId)
	if err != nil {
		return nil, ParseError(err).Err()
	}

	return MapDocToResponse(app), nil
}

func (s *Server) CreateApp(ctx context.Context, req *protos.CreateAppRequest) (*protos.App, error) {
	userId, err := GetAuthUser(ctx)
	if err != nil {
		return nil, ParseError(err).Err()
	}

	payload := &domain.AppPayload{
		OwnerID:     userId,
		Name:        req.App.Name,
		Description: req.App.Description,
		Environment: req.App.Environment.String(),
		OwnerType:   req.App.OwnerType.String(),
	}

	app, err := s.Service.CreateApp(payload)
	if err != nil {
		return nil, ParseError(err).Err()
	}

	return MapDocToResponse(app), nil
}

func (s *Server) RefreshAppKey(ctx context.Context, req *protos.RefreshAppKeyRequest) (*protos.RefreshAppKeyResponse, error) {
	userId, err := GetAuthUser(ctx)
	if err != nil {
		return nil, ParseError(err).Err()
	}

	newKey, err := s.Service.RefreshKey(userId, req.AppId)
	if err != nil {
		return nil, ParseError(err).Err()
	}

	return &protos.RefreshAppKeyResponse{NewKey: newKey}, nil
}

func (s *Server) DecodeAppKey(ctx context.Context, req *protos.DecodeAppKeyRequest) (*protos.DecodeAppKeyResponse, error) {
	appId, err := s.Service.DecodeKey(req.AppKey)
	if err != nil {
		return nil, ParseError(err).Err()
	}

	return &protos.DecodeAppKeyResponse{AppId: appId}, nil
}

func (s *Server) DeleteApp(ctx context.Context, req *protos.DeleteAppRequest) (*emptypb.Empty, error) {
	userId, err := GetAuthUser(ctx)
	if err != nil {
		return nil, ParseError(err).Err()
	}

	if err := s.Service.DeleteApp(userId, req.AppId); err != nil {
		return nil, ParseError(err).Err()
	}

	return &emptypb.Empty{}, nil
}
