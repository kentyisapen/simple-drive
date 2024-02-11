// internal/interfaces/grpc/server/file_manager_server.go
package server

import (
	"context"

	"github.com/kentyisapen/simple-drive/internal/application/service"
	"github.com/kentyisapen/simple-drive/pb"
)

type FileManagerServer struct {
	pb.UnimplementedFileManagerServer
	healthCheckService *service.HealthCheckService
	// 他のサービスやユースケースをフィールドとして追加
}

func NewFileManagerServer(healthCheckService *service.HealthCheckService /* 他の依存サービス */) *FileManagerServer {
	return &FileManagerServer{
		healthCheckService: healthCheckService,
		// 他のフィールドを初期化
	}
}

func (s *FileManagerServer) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	status := s.healthCheckService.Check()
	return &pb.HealthCheckResponse{
		Status: status,
	}, nil
}

// 他のメソッドをFileManagerServerに追加
func (s *FileManagerServer) ItemCreate(ctx context.Context, req *pb.ItemCreateRequest) (*pb.Item, error) {
	// 他のアプリケーションサービスやユースケースを使用するロジック
	return &pb.Item{
		// 応答内容
	}, nil
}
