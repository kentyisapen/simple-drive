// internal/domain/repository/item_repository.go
package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/kentyisapen/simple-drive/internal/domain/model"
)

type ItemPostgresRepository interface {
	CreateItem(ctx context.Context, item model.Item) (model.Item, error)
	ListItems(ctx context.Context, parentID string, page, size int32) ([]model.Item, error)
	CountItems(ctx context.Context, parentID string) (int64, error)
}

type ItemMinioRepository interface {
	SaveContent(ctx context.Context, itemID uuid.UUID, content []byte) (string, error)
	DeleteContent(ctx context.Context, itemID string) error
}
