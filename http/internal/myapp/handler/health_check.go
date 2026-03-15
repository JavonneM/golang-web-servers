package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (sh *SystemHandlerImpl) HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Process input
	response, err := sh.SystemService.HealthCheck()
	responseAsBytes, err := json.Marshal(response)
	if err != nil {
		/// fail
		fmt.Println("failed to marshal response")
		w.WriteHeader(500)
		return
	}
	w.Write(responseAsBytes)
}
