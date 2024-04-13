// internal/application/service/item_create_service.go
package appservice

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/kentyisapen/simple-drive/internal/application/usecase"
	"github.com/kentyisapen/simple-drive/internal/domain/model"
	"github.com/kentyisapen/simple-drive/pb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type ItemCreateService struct {
	itemCreateUsecase *usecase.ItemCreateUsecase
}

func NewItemCreateService(itemCreateUsecase *usecase.ItemCreateUsecase) *ItemCreateService {
	return &ItemCreateService{
		itemCreateUsecase: itemCreateUsecase,
	}
}

func (s *ItemCreateService) Execute(ctx context.Context, req *pb.ItemCreateRequest) (model.Item, error) {
	if req == nil {
		return model.Item{}, fmt.Errorf("request is nil")
	}

	if req.GetParentId() == nil || req.GetParentId().Value == "" {
		return model.Item{}, fmt.Errorf("parent ID is required")
	}

	parentUuid, err := parseUuid(req.GetParentId())
	if err != nil {
		return model.Item{}, fmt.Errorf("invalid parent UUID: %v", err)
	}

	return s.itemCreateUsecase.Execute(ctx, req, parentUuid)
}

func parseUuid(uuidValue *wrapperspb.StringValue) (*uuid.UUID, error) {
	parsedUuid, err := uuid.Parse(uuidValue.Value)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %v", err)
	}
	return &parsedUuid, nil
}
