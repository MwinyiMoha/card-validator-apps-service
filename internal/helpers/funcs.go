package helpers

import (
	"context"
	"time"
)

func NewTimeoutContext(t int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
}
