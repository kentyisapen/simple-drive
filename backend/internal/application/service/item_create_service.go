// internal/application/service/item_create_service.go
package appservice

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/kentyisapen/simple-drive/internal/application/usecase"
	"github.com/kentyisapen/simple-drive/internal/domain/model"
	"github.com/kentyisapen/simple-drive/pb"
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

	var parentUuid = uuid.Nil
	parsed, err := uuid.Parse(req.GetParentId())
	if err != nil {
		parentUuid = parsed
	}

	return s.itemCreateUsecase.Execute(ctx, req, parentUuid)
}
