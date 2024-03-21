package usecase

import (
	"context"

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
	var tempUuid *uuid.UUID = nil
	if parsedUuid, err := uuid.Parse(req.GetParentId().Value); err == nil {
		tempUuid = &parsedUuid
	}

	itemID := uuid.New()
	if err := icu.minioRepo.SaveContent(ctx, itemID, req.GetFile()); err != nil {
		return model.Item{}, err
	}

	size := int64(len(req.GetFile()))
	item := model.Item{
		ID:       itemID,
		ParentID: tempUuid,
		Name:     req.GetName(),
		Type:     model.FileType,
		Size:     &size, // int64の値をポインタに変換
	}

	savedItem, err := icu.postgresRepo.CreateItem(ctx, item)
	if err != nil {
		return model.Item{}, err
	}
	return savedItem, nil
}
