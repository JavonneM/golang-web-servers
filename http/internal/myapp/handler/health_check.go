package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/javonnem/web_server/http/internal/myapp/dto"
	"github.com/javonnem/web_server/http/internal/myapp/service"
	"github.com/javonnem/web_server/http/pkg/server"
)

type SystemHandlerImpl struct {
	SystemService service.SystemService
}

type SystemHandler interface {
	HealthCheck(http.ResponseWriter, *http.Request)
	TestApiHandler(*http.Request, dto.TestApiHandlerRequest) server.HandlerResponse
	TestApiHandlerWithPathValues(r *http.Request, request any) server.HandlerResponse
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

func (sh *SystemHandlerImpl) TestApiHandler(r *http.Request, request dto.TestApiHandlerRequest) server.HandlerResponse {
	// Process input
	fmt.Printf("%+v\n", request)
	fmt.Println(request.SomeField)
	// Do work
	response, err := sh.SystemService.HealthCheck()
	_, err = json.Marshal(response)
	if err != nil {
		return server.HandlerResponse{
			ProtocolResponseCode: 500,
			Data: server.HttpError{
				Error: fmt.Errorf("failed to marshal response %w", err),
			},
		}
	}
	return server.HandlerResponse{
		ProtocolResponseCode: 200,
		Data:                 response,
	}
}

func (sh *SystemHandlerImpl) TestApiHandlerWithPathValues(r *http.Request, request any) server.HandlerResponse {
	// Process input
	id := r.PathValue("id")
	pageAsString := r.URL.Query().Get("page")
	fmt.Printf("id %s\n", id)
	fmt.Printf("page %s\n", pageAsString)
	// Do work
	response, err := sh.SystemService.HealthCheck()
	_, err = json.Marshal(response)
	if err != nil {
		return server.HandlerResponse{
			ProtocolResponseCode: 500,
			Data: server.HttpError{
				Error: fmt.Errorf("failed to marshal response %w", err),
			},
		}
	}
	return server.HandlerResponse{
		ProtocolResponseCode: 200,
		Data:                 response,
	}
}
