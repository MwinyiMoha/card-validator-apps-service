package repository

import "context"

func (r *Repository) Disconnect(ctx context.Context) error {
	return r.Client.Disconnect((ctx))
}
