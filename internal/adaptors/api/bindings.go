package api

import (
	"card-validator-apps-service/internal/core/domain"
	protos "card-validator-apps-service/internal/gen"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapDocToResponse(doc *domain.App) *protos.App {
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

func ParseError(err error) error {
	fmt.Println(err)
	return err
}
