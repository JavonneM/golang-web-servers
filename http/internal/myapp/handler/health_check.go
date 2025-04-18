package handler

import (
	"time"

	"github.com/javonnem/web_server/http/internal/myapp/dto"
)

func HealthCheck(payload any) (any, error) {
	time.Sleep(5 * time.Second)
	return dto.HealthCheckResponse{
		SomeText: "success",
	}, nil
}
