// backend/internal/application/usecase/item_create.go
package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kentyisapen/simple-drive/internal/domain/model"
	"github.com/kentyisapen/simple-drive/internal/infrastructures/repository" // repositoryパッケージをインポート
	"github.com/kentyisapen/simple-drive/pb"
)

type ItemCreateUsecase struct {
	postgresRepo repository.ItemPostgresRepository // ItemPostgresRepositoryを使用
	minioRepo    repository.ItemMinioRepository    // ItemMinioRepositoryを使用
}

func NewItemCreateUsecase(postgresRepo repository.ItemPostgresRepository, minioRepo repository.ItemMinioRepository) *ItemCreateUsecase {
	return &ItemCreateUsecase{
		postgresRepo: postgresRepo,
		minioRepo:    minioRepo,
	}
}

func (icu *ItemCreateUsecase) Execute(ctx context.Context, req *pb.ItemCreateRequest) (model.Item, error) {
	if icu.postgresRepo == nil {
		return model.Item{}, fmt.Errorf("postgres repository is nil")
	}

	if icu.minioRepo == nil {
		return model.Item{}, fmt.Errorf("minio repository is nil")
	}

	if req == nil {
		return model.Item{}, fmt.Errorf("request is nil")
	}

	var tempUuid *uuid.UUID = nil
	if req.GetParentId() != nil && req.GetParentId().Value != "" {
		parsedUuid, err := uuid.Parse(req.GetParentId().Value)
		if err == nil {
			tempUuid = &parsedUuid
		} else {
			// エラー処理（無効なUUID形式が渡された場合）
		}
	}

	// ファイルの内容を保存する前に、req.GetFile() が nil でないことを確認します。
	if req.GetFile() == nil {
		return model.Item{}, fmt.Errorf("file content is nil")
	}

	itemID := uuid.New()
	objectKey, err := icu.minioRepo.SaveContent(ctx, itemID, req.GetFile())
	if err != nil {
		return model.Item{}, err
	}

	now := time.Now()
	size := int64(len(req.GetFile()))
	item := model.Item{
		ID:             itemID,
		ParentID:       tempUuid,
		Name:           req.GetName(),
		Type:           model.FileType,
		Size:           &size,
		CreatedAt:      now,
		LastModifiedAt: now,
		MinioObjectKey: objectKey, // MinIOから返されたオブジェクトキー
	}

	savedItem, err := icu.postgresRepo.CreateItem(ctx, item)
	if err != nil {
		return model.Item{}, err
	}
	return savedItem, nil
}
