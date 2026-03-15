package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/javonnem/web_server/http/internal/myapp/dto"
	"github.com/javonnem/web_server/http/internal/myapp/service"
)

type SystemHandlerImpl struct {
	SystemService service.SystemService
}

type SystemHandler interface {
	HealthCheck(http.ResponseWriter, *http.Request)
}

func NewSystemHandler(s service.SystemService) SystemHandler {
	return &SystemHandlerImpl{
		SystemService: s,
	}
}

func (*SystemHandlerImpl) HealthCheck(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)

	response := &dto.HealthCheckResponse{
		ServerStatusSuccess:   true,
		DatabaseStatusSuccess: false,
	}
	responseAsBytes, err := json.Marshal(response)
	if err != nil {
		/// fail
		fmt.Println("failed to marshal response")
	}
	w.Write(responseAsBytes)
}
