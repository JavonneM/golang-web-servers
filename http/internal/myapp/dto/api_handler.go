package dto


type TestApiHandlerRequest struct {
	SomeField string `json:"somefield"`	
}
type TestApiHandlerResponse struct {
	ServerStatusSuccess   bool `json:"server_status_success"`
	DatabaseStatusSuccess bool `json:"database_status_success"`
}

