// internal/application/usecase/item_list.go
package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/kentyisapen/simple-drive/internal/domain/repository"
	"github.com/kentyisapen/simple-drive/pb"
)

type ItemListUsecase struct {
	postgresRepo repository.ItemPostgresRepository
}

func NewItemListUsecase(postgresRepo repository.ItemPostgresRepository) *ItemListUsecase {
	return &ItemListUsecase{
		postgresRepo: postgresRepo,
	}
}

func (ilu *ItemListUsecase) Execute(ctx context.Context, parentId uuid.UUID, page, size int32) (*pb.ListItemsResponse, error) {
	if page <= 0 {
		page = 1
	}

	if size <= 0 {
		size = 20
	}

	items, err := ilu.postgresRepo.ListItems(ctx, parentId.String(), page, size)
	if err != nil {
		return nil, err
	}

	totalItems, err := ilu.postgresRepo.CountItems(ctx, parentId.String())
	if err != nil {
		return nil, err
	}

	totalPages := totalItems / int64(size)
	if totalItems%int64(size) != 0 {
		totalPages++
	}

	pbItems := make([]*pb.Item, len(items))
	for i, item := range items {
		pbItems[i] = item.ToPB()
	}

	response := &pb.ListItemsResponse{
		Items: pbItems,
		Paging: &pb.Paging{
			CurrentPage: int32(page),
			PageSize:    size,
			TotalItems:  totalItems,
			TotalPages:  int32(totalPages),
		},
	}

	return response, nil
}
