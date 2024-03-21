// backend/internal/infrastructures/repository/minio/item_minio_repository.go
package minio

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/kentyisapen/simple-drive/internal/infrastructures/repository"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type itemMinioRepository struct {
	minioClient *minio.Client
}

func NewItemMinioRepository(endpoint, accessKeyID, secretAccessKey string) (repository.ItemMinioRepository, error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: true, // SSLを使用するかどうか
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %w", err)
	}

	return &itemMinioRepository{minioClient: minioClient}, nil
}

func (r *itemMinioRepository) SaveContent(ctx context.Context, itemID uuid.UUID, content []byte) error {
	reader := bytes.NewReader(content)
	contentType := http.DetectContentType(content)
	size := int64(len(content))
	_, err := r.minioClient.PutObject(ctx, os.Getenv("MINIO_BUCKET_NAME"), itemID.String(), reader, size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return fmt.Errorf("error while saving content to MinIO: %w", err)
	}
	return nil
}
