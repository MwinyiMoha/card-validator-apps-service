package api

import (
	"card-validator-apps-service/internal/core/domain"
	protos "card-validator-apps-service/internal/gen"

	"github.com/mwinyimoha/card-validator-utils/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func appDocToProto(doc *domain.App) *protos.App {
	return &protos.App{
		Id:          doc.ID.Hex(),
		Name:        doc.Name,
		Description: doc.Description,
		Environment: protos.Environment(protos.Environment_value[doc.Environment]),
		OwnerType:   protos.OwnerType(protos.OwnerType_value[doc.OwnerType]),
		AppKey:      doc.AppKey,
		CreatedAt:   timestamppb.New(doc.CreatedAt),
		LastUpdated: timestamppb.New(doc.LastUpdated),
	}
}

func parseError(err error) *status.Status {
	if cerr, ok := err.(*errors.Error); ok {
		return status.Convert(cerr)
	}

	if verr, ok := err.(*errors.ValidationError); ok {
		return status.Convert(verr)
	}

	return status.New(codes.Internal, err.Error())
}
