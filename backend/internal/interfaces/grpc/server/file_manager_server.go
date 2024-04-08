// internal/interfaces/grpc/server/file_manager_server.go
package server

import (
	"context"

	appservice "github.com/kentyisapen/simple-drive/internal/application/service"
	"github.com/kentyisapen/simple-drive/pb"
)

type FileManagerServer struct {
	pb.UnimplementedFileManagerServer
	healthCheckService *appservice.HealthCheckService
	itemCreateService  *appservice.ItemCreateService
	itemListService    *appservice.ItemListService
}

func NewFileManagerServer(
	healthCheckService *appservice.HealthCheckService,
	itemCreateService *appservice.ItemCreateService,
	itemListService *appservice.ItemListService,
) *FileManagerServer {
	return &FileManagerServer{
		healthCheckService: healthCheckService,
		itemCreateService:  itemCreateService,
		itemListService:    itemListService,
	}
}

func (s *FileManagerServer) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	status := s.healthCheckService.Check()
	return &pb.HealthCheckResponse{
		Status: status,
	}, nil
}

func (s *FileManagerServer) CreateItem(ctx context.Context, req *pb.ItemCreateRequest) (*pb.Item, error) {
	// 他のアプリケーションサービスやユースケースを使用するロジック
	item, err := s.itemCreateService.Execute(ctx, req)
	if err != nil {
		return &pb.Item{}, err
	}
	return item.ToPB(), nil
}

func (s *FileManagerServer) ListItems(ctx context.Context, req *pb.ListItemsRequest) (*pb.ListItemsResponse, error) {
	response, err := s.itemListService.Execute(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
