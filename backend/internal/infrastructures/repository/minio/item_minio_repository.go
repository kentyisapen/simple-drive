// internal/infrastructures/repository/minio/item_minio_repository.go
package minio

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/kentyisapen/simple-drive/internal/domain/repository"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type itemMinioRepository struct {
	minioClient *minio.Client
}

func NewItemMinioRepository(endpoint, accessKeyID, secretAccessKey string) (repository.ItemMinioRepository, error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false, // SSLを使用するかどうか
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %w", err)
	}

	return &itemMinioRepository{minioClient: minioClient}, nil
}

func (r *itemMinioRepository) SaveContent(ctx context.Context, itemID uuid.UUID, content []byte) (string, error) {
	reader := bytes.NewReader(content)
	contentType := http.DetectContentType(content)
	size := int64(len(content))
	info, err := r.minioClient.PutObject(ctx, os.Getenv("MINIO_BUCKET_NAME"), itemID.String(), reader, size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", fmt.Errorf("error while saving content to MinIO: %w", err)
	}
	return info.Key, nil
}

func (r *itemMinioRepository) DeleteContent(ctx context.Context, itemID string) error {
	err := r.minioClient.RemoveObject(ctx, os.Getenv("MINIO_BUCKET_NAME"), itemID, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("error while deleting content from MinIO: %w", err)
	}
	return nil
}
