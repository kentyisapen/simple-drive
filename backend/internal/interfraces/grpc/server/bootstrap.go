// internal/interfaces/grpc/server/bootstrap.go
package server

import (
	"github.com/kentyisapen/simple-drive/internal/application/service"
	"github.com/kentyisapen/simple-drive/internal/application/usecase"
	pb "github.com/kentyisapen/simple-drive/pb"
	"google.golang.org/grpc"
)

// BootstrapServices は、gRPCサーバーにサービスを登録し、依存性を注入します。
func BootstrapServices(grpcServer *grpc.Server) {
	// リポジトリの初期化
	// itemRepository := repository.NewItemRepository()

	// ユースケースの初期化
	healthCheckUsecase := usecase.NewHealthCheckUsecase()
	healthCheckService := service.NewHealthCheckService(healthCheckUsecase)

	// anyUsecase := usecase.XXX(YyyRepository)

	// 他のサービスやユースケースのインスタンスを作成

	// FileManagerServerのインスタンスを作成し、すべての依存性を注入
	fileManagerServer := NewFileManagerServer(healthCheckService /*, 他のサービスやユースケースのインスタンス */)
	pb.RegisterFileManagerServer(grpcServer, fileManagerServer)
}
