// backend/internal/infrastructures/repository/item_repository.go
package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/kentyisapen/simple-drive/internal/domain/model"
)

type ItemPostgresRepository interface {
	CreateItem(ctx context.Context, item model.Item) (model.Item, error)
}

type ItemMinioRepository interface {
	SaveContent(ctx context.Context, itemID uuid.UUID, content []byte) error
}
