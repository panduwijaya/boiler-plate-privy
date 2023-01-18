package repositories

import (
	"context"
)

// Storer store contract
type Uploader interface {
	Store(ctx context.Context, param interface{}) (int64, error)
}
