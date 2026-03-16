/// Service Package
package service

import (
	"database/sql"

	"github.com/javonnem/web_server/http/internal/myapp/dto"
)

type SystemServiceImpl struct {
	dbConn *sql.DB
}

type SystemService interface {
	HealthCheck() (*dto.HealthCheckResponse, error)
}

func NewSystemService(dbConn *sql.DB) SystemService {
	return &SystemServiceImpl{
		dbConn: dbConn,
	}
}

func (ss *SystemServiceImpl) HealthCheck() (
	*dto.HealthCheckResponse, 
	error,
) {
	err := ss.dbConn.Ping()

	return &dto.HealthCheckResponse{
		ServerStatusSuccess:   true,
		DatabaseStatusSuccess: err == nil,
	}, nil

}

