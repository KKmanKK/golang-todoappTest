package user_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/KKmanKK/golang-todoappTest/internal/core/logger"
	core_http_response "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/response"
	core_http_utils "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/utils"
)

type GetUsersResponse []UserDTOResponse

func (h *UsersHTTPHander) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffset(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit or offset query param")
		return
	}
	log.Debug("Get Users handler")

	usersDomains, err := h.usersService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get users")
		return
	}

	response := GetUsersResponse(usersDTOFromDomains(usersDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getLimitOffset(r *http.Request) (*int, *int, error) {
	limit, err := core_http_utils.GetIntQuertyParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}
	offset, err := core_http_utils.GetIntQuertyParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}
	return limit, offset, nil
}
