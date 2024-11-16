package api

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/mwinyimoha/card-validator-utils/errors"
)

const userIdHeader = "X-User-ID"

func AuthFn(ctx context.Context) (context.Context, error) {
	userId := metadata.ExtractIncoming(ctx).Get(userIdHeader)
	if userId == "" {
		return nil, errors.NewErrorf(errors.Unauthenticated, "request is not authenticated")
	}

	return ctx, nil
}

func GetAuthUser(ctx context.Context) (string, error) {
	userId := metadata.ExtractIncoming(ctx).Get(userIdHeader)
	if userId == "" {
		return "", errors.NewErrorf(errors.Unauthenticated, "request is not authenticated")
	}

	return userId, nil
}

func SkipAuth(_ context.Context, c interceptors.CallMeta) bool {
	return c.FullMethod() != "/protos.AppsService/DecodeAppKey"
}
