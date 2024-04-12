// internal/application/usecase/item_create.go
package usecase

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/kentyisapen/simple-drive/internal/domain/model"
	"github.com/kentyisapen/simple-drive/internal/domain/repository"

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
	_, err := icu.minioRepo.SaveContent(ctx, itemID, req.GetFile())
	if err != nil {
		return model.Item{}, err
	}

	thumbnailID := uuid.New()
	_, err = createThumbnailFromBinary(thumbnailID, req.GetFile())
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
	}

	savedItem, err := icu.postgresRepo.CreateItem(ctx, item)
	if err != nil {
		deleteErr := icu.minioRepo.DeleteContent(ctx, itemID.String())
		if deleteErr != nil {
			return model.Item{}, err
		}
		return model.Item{}, err
	}
	return savedItem, nil
}

func createThumbnailFromBinary(thumbnailID uuid.UUID, data []byte) (uuid.UUID, error) {
	mimeType := http.DetectContentType(data)
	var nullUUID uuid.NullUUID
	if mimeType != "application/pdf" && !strings.HasPrefix(mimeType, "image/") {
		return nullUUID.UUID, fmt.Errorf("unsupported file type: %s", mimeType)
	}

	tmpFile, err := os.CreateTemp("", "source")
	if err != nil {
		return nullUUID.UUID, err
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(data); err != nil {
		return nullUUID.UUID, err
	}

	outputFileName := thumbnailID.String() + ".webp"

	outputFilePath := "tmp/" + outputFileName

	inputFilePath := tmpFile.Name()
	if mimeType == "application/pdf" {
		inputFilePath += "[0]"
	}

	cmd := exec.Command("convert", inputFilePath, "-resize", "420x", "-background", "white", "-alpha", "remove", outputFilePath)
	if err := cmd.Run(); err != nil {
		return nullUUID.UUID, err
	}

	return thumbnailID, nil
}
