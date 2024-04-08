// internal/application/service/item_create_service.go
package appservice

import (
	"context"

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
	return s.itemCreateUsecase.Execute(ctx, req)
}
