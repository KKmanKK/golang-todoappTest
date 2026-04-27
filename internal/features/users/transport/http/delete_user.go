package user_transport_http

import (
	"net/http"

	core_logger "github.com/KKmanKK/golang-todoappTest/internal/core/logger"
	core_http_response "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/response"
	core_http_utils "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/utils"
)

func (h *UsersHTTPHander) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userId, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userId path value")
		return
	}

	if err = h.usersService.DeleteUser(ctx, userId); err != nil {
		responseHandler.ErrorResponse(err, "falied to delete user")
		return
	}

	responseHandler.NoContentResponse()
}
