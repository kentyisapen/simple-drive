// internal/application/service/item_list_service.go
package appservice

import (
	"context"

	"github.com/google/uuid"
	"github.com/kentyisapen/simple-drive/internal/application/usecase"
	"github.com/kentyisapen/simple-drive/pb"
)

type ItemListService struct {
	itemListUsecase *usecase.ItemListUsecase
}

func NewItemListService(itemListUsecase *usecase.ItemListUsecase) *ItemListService {
	return &ItemListService{
		itemListUsecase: itemListUsecase,
	}
}

func (s *ItemListService) Execute(ctx context.Context, req *pb.ListItemsRequest) (*pb.ListItemsResponse, error) {

	var parentId = uuid.Nil
	parsed, err := uuid.Parse(req.GetParentId())
	if err == nil {
		parentId = parsed
	}

	response, err := s.itemListUsecase.Execute(ctx, parentId, req.GetPage(), req.GetSize())
	if err != nil {
		return nil, err
	}

	return response, nil
}
