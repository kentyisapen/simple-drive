// internal/application/service/health_check_service.go
package appservice

import (
	"github.com/kentyisapen/simple-drive/internal/application/usecase"
)

type HealthCheckService struct {
	healthCheckUsecase *usecase.HealthCheckUsecase
}

func NewHealthCheckService(healthCheckUsecase *usecase.HealthCheckUsecase) *HealthCheckService {
	return &HealthCheckService{
		healthCheckUsecase: healthCheckUsecase,
	}
}

func (s *HealthCheckService) Check() string {
	return s.healthCheckUsecase.Check()
}
