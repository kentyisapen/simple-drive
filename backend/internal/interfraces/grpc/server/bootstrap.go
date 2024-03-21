package server

import (
	"fmt"
	"os"

	appservice "github.com/kentyisapen/simple-drive/internal/application/service"
	"github.com/kentyisapen/simple-drive/internal/application/usecase"
	"github.com/kentyisapen/simple-drive/internal/infrastructures/repository/minio"
	"github.com/kentyisapen/simple-drive/internal/infrastructures/repository/postgresql"
	pb "github.com/kentyisapen/simple-drive/pb"
	"google.golang.org/grpc"
)

// BootstrapServices は、gRPCサーバーにサービスを登録し、依存性を注入します。
func BootstrapServices(grpcServer *grpc.Server) error {
	// 環境変数から接続詳細を取得
	dbConnectionString := os.Getenv("DATABASE_URL")
	minioEndpoint := os.Getenv("MINIO_ENDPOINT")
	minioAccessKey := os.Getenv("MINIO_ACCESS_KEY")
	minioSecretKey := os.Getenv("MINIO_SECRET_KEY")

	// PostgreSQLリポジトリの初期化
	itemPostgresRepository, err := postgresql.NewItemPostgresqlRepository(dbConnectionString)
	if err != nil {
		fmt.Println("PostgreSQL repository initialization failed:", err)
		return err
	}

	// MinIOリポジトリの初期化
	itemMinioRepository, err := minio.NewItemMinioRepository(minioEndpoint, minioAccessKey, minioSecretKey)
	if err != nil {
		fmt.Println("MinIO repository initialization failed:", err)
		return err
	}

	// ユースケースの初期化
	healthCheckUsecase := usecase.NewHealthCheckUsecase()
	healthCheckService := appservice.NewHealthCheckService(healthCheckUsecase)

	itemCreateUsecase := usecase.NewItemCreateUsecase(itemPostgresRepository, itemMinioRepository)
	itemCreateService := appservice.NewItemCreateService(itemCreateUsecase)

	// FileManagerServerのインスタンスを作成し、すべての依存性を注入
	fileManagerServer := NewFileManagerServer(healthCheckService, itemCreateService)
	pb.RegisterFileManagerServer(grpcServer, fileManagerServer)

	return nil
}
