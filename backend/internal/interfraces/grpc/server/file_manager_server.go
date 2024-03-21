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
	// 他のサービスやユースケースをフィールドとして追加
}

func NewFileManagerServer(healthCheckService *appservice.HealthCheckService, itemCreateService *appservice.ItemCreateService) *FileManagerServer {
	return &FileManagerServer{
		healthCheckService: healthCheckService,
		itemCreateService:  itemCreateService,
	}
}

func (s *FileManagerServer) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	status := s.healthCheckService.Check()
	return &pb.HealthCheckResponse{
		Status: status,
	}, nil
}

// 他のメソッドをFileManagerServerに追加
func (s *FileManagerServer) CreateItem(ctx context.Context, req *pb.ItemCreateRequest) (*pb.Item, error) {
	// 他のアプリケーションサービスやユースケースを使用するロジック
	item, err := s.itemCreateService.Execute(ctx, req)
	if err != nil {
		return &pb.Item{}, err
	}
	return item.ToPB(), nil
}
