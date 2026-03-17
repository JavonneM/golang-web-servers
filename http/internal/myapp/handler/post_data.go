package handler

// import (
// 	"encoding/json"
// 	"net/http"
// 	"time"
// 
// 	"github.com/javonnem/web_server/http/internal/myapp/dto"
// 	"github.com/javonnem/web_server/http/pkg/server"
// 	"github.com/javonnem/web_server/http/pkg/util"
// )
// 
// func PostDataWithPathParameters(ctx util.AppContext, r *http.Request) (any, error) {
// 	id := r.PathValue("id")
// 	var body dto.PostDataRequest
// 
// 	err := json.NewDecoder(r.Body).Decode(&body)
// 	if err != nil {
// 		return nil, server.ErrHttpBadRequestError
// 	}
// 
// 	return dto.PostDataResponse{
// 		TimeStamp: time.Now().UTC(),
// 		PathId: id,
// 		Data: body.Data,
// 	}, nil
// }
