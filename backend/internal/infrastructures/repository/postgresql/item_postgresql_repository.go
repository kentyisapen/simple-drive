// backend/internal/infrastructures/repository/postgresql/item_postgresql_repository.go
package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kentyisapen/simple-drive/internal/domain/model"
	"github.com/kentyisapen/simple-drive/internal/infrastructures/repository"
)

type itemPostgresqlRepository struct {
	db *sql.DB
}

func NewItemPostgresqlRepository(connectionString string) (repository.ItemPostgresRepository, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &itemPostgresqlRepository{db: db}, nil
}

func (r *itemPostgresqlRepository) CreateItem(ctx context.Context, item model.Item) (model.Item, error) {
	query := `INSERT INTO items (id, parent_id, name, type, size, created_at, last_modified_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query, item.ID, item.ParentID, item.Name, item.Type, item.Size, item.CreatedAt, item.LastModifiedAt)
	if err != nil {
		return model.Item{}, fmt.Errorf("error while inserting item into db: %w", err)
	}
	return item, nil
}
