// internal/application/service/item_list_service.go
package appservice

import (
	"context"

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

	var parentID *string
	if req.GetParentId() != nil {
		parentIDVal := req.GetParentId().GetValue()
		parentID = &parentIDVal
	}

	response, err := s.itemListUsecase.Execute(ctx, parentID, req.GetPage(), req.GetSize())
	if err != nil {
		return nil, err
	}

	return response, nil
}
