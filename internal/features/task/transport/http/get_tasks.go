package task_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/KKmanKK/golang-todoappTest/internal/core/logger"
	core_http_request "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/request"
	core_http_response "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/response"
)

type GetTasksResponse []TaskDTOResponse

func (h *TaskHTTPHandler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userId, limit, offset, err := getUserIdLimitOffset(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user_id/limit/offset query param")
		return
	}

	tasks, err := h.taskService.GetTasks(ctx, userId, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "falied get tasks")
		return
	}

	response := GetTasksResponse(tasksDTOFromDomains(tasks...))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getUserIdLimitOffset(r *http.Request) (*int, *int, *int, error) {
	const (
		userIdQueryParamKey = "user_id"
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)
	userId, err := core_http_request.GetIntQuertyParam(r, userIdQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}
	limit, err := core_http_request.GetIntQuertyParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}
	offset, err := core_http_request.GetIntQuertyParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}
	return userId, limit, offset, nil
}
