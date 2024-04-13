// internal/application/usecase/item_create.go

package usecase

import (
	"context"
	"log"
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
	postgresRepo repository.ItemPostgresRepository
	minioRepo    repository.ItemMinioRepository
}

func NewItemCreateUsecase(postgresRepo repository.ItemPostgresRepository, minioRepo repository.ItemMinioRepository) *ItemCreateUsecase {
	return &ItemCreateUsecase{
		postgresRepo: postgresRepo,
		minioRepo:    minioRepo,
	}
}

func (icu *ItemCreateUsecase) Execute(ctx context.Context, req *pb.ItemCreateRequest, parentUuid *uuid.UUID) (model.Item, error) {
	itemID, thumbnailID, savedItem, err := icu.handleItemAndThumbnailCreation(ctx, req, parentUuid)
	if err != nil {
		icu.cleanupOnError(ctx, itemID.String(), thumbnailID.String())
		return model.Item{}, err
	}

	return savedItem, nil
}

func (icu *ItemCreateUsecase) handleItemAndThumbnailCreation(ctx context.Context, req *pb.ItemCreateRequest, parentUuid *uuid.UUID) (uuid.UUID, uuid.UUID, model.Item, error) {
	itemID := uuid.New()
	var nullUUID uuid.NullUUID
	if err := icu.saveFileContent(ctx, itemID, req.GetFile()); err != nil {
		return itemID, nullUUID.UUID, model.Item{}, err
	}

	thumbnailID, err := icu.createThumbnailFromBinary(ctx, req.GetFile())
	if err != nil {
		return itemID, thumbnailID, model.Item{}, err
	}

	now := time.Now()
	size := int64(len(req.GetFile()))
	item := model.Item{
		ID:             itemID,
		ParentID:       parentUuid,
		ThumbnailId:    thumbnailID,
		Name:           req.GetName(),
		Type:           model.FileType,
		Size:           &size,
		CreatedAt:      now,
		LastModifiedAt: now,
	}

	savedItem, err := icu.saveItemToDatabase(ctx, item)
	if err != nil {
		return itemID, thumbnailID, model.Item{}, err
	}

	return itemID, thumbnailID, savedItem, nil
}

func (icu *ItemCreateUsecase) createThumbnailFromBinary(ctx context.Context, data []byte) (uuid.UUID, error) {
	var thumbnailID = uuid.New()
	mimeType := http.DetectContentType(data)
	if mimeType != "application/pdf" && !strings.HasPrefix(mimeType, "image/") {
		return uuid.Nil, nil
	}

	tmpFile, err := os.CreateTemp("", "source")
	if err != nil {
		return uuid.Nil, err
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(data); err != nil {
		return uuid.Nil, err
	}

	outputFileName := thumbnailID.String() + ".webp"
	outputFilePath := "tmp/" + outputFileName
	inputFilePath := tmpFile.Name()
	if mimeType == "application/pdf" {
		inputFilePath += "[0]"
	}

	cmd := exec.Command("convert", inputFilePath, "-resize", "420x", "-background", "white", "-alpha", "remove", outputFilePath)
	if err := cmd.Run(); err != nil {
		return uuid.Nil, err
	}
	defer os.Remove(outputFilePath)

	bin, err := os.ReadFile(outputFilePath)
	if err != nil {
		return uuid.Nil, err
	}
	_, err = icu.minioRepo.SaveContent(ctx, thumbnailID, bin)
	if err != nil {
		return uuid.Nil, err
	}

	return thumbnailID, nil
}

func (icu *ItemCreateUsecase) saveFileContent(ctx context.Context, itemID uuid.UUID, fileContent []byte) error {
	_, err := icu.minioRepo.SaveContent(ctx, itemID, fileContent)
	return err
}

func (icu *ItemCreateUsecase) saveItemToDatabase(ctx context.Context, item model.Item) (model.Item, error) {
	return icu.postgresRepo.CreateItem(ctx, item)
}

func (icu *ItemCreateUsecase) cleanupOnError(ctx context.Context, itemID string, thumbnailID string) {
	if delErr := icu.minioRepo.DeleteContent(ctx, itemID); delErr != nil {
		log.Fatal(delErr)
	}
	if delErr := icu.minioRepo.DeleteContent(ctx, thumbnailID); delErr != nil {
		log.Fatal(delErr)
	}
}
