// internal/infrastructures/repository/postgresql/item_postgresql_repository.go
package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kentyisapen/simple-drive/internal/domain/model"
	"github.com/kentyisapen/simple-drive/internal/domain/repository"
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
	query := `INSERT INTO items (id, parent_id, name, type, size, created_at, last_modified_at, thumbnail_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.ExecContext(ctx, query, item.ID, item.ParentID, item.Name, item.Type, item.Size, item.CreatedAt, item.LastModifiedAt, item.ThumbnailId)
	if err != nil {
		return model.Item{}, fmt.Errorf("error while inserting item into db: %w", err)
	}
	return item, nil
}

func (r *itemPostgresqlRepository) ListItems(ctx context.Context, parentID string, page, size int32) ([]model.Item, error) {
	var items []model.Item
	var rows *sql.Rows
	var err error

	query := `SELECT id, parent_id, name, type, size, created_at, last_modified_at, thumbnail_id FROM items WHERE parent_id = $1 LIMIT $2 OFFSET $3`
	rows, err = r.db.QueryContext(ctx, query, parentID, size, (page-1)*size)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.Item
		if err := rows.Scan(&item.ID, &item.ParentID, &item.Name, &item.Type, &item.Size, &item.CreatedAt, &item.LastModifiedAt, &item.ThumbnailId); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *itemPostgresqlRepository) CountItems(ctx context.Context, parentID string) (int64, error) {
	var count int64
	var query string
	var err error // errを関数のスコープで定義する

	query = "SELECT COUNT(*) FROM items WHERE parent_id = $1"
	err = r.db.QueryRowContext(ctx, query, parentID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
