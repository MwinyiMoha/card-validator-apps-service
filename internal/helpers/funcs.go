package helpers

import (
	"context"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/mwinyimoha/card-validator-utils/errors"
)

var randReader = rand.Read

func NewTimeoutContext(t int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
}

func NewRandomCode(length int) (string, error) {
	code := make([]byte, length)

	_, err := randReader(code)
	if err != nil {
		return "", errors.WrapError(err, errors.Internal, "could not generate random string")
	}

	return fmt.Sprintf("%x", code), nil
}
