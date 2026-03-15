package dto

type HealthCheckResponse struct {
	ServerStatusSuccess   bool `json:"server_status_success"`
	DatabaseStatusSuccess bool `json:"database_status_success"`
}
