package task_transport_http

import (
	"net/http"

	core_logger "github.com/KKmanKK/golang-todoappTest/internal/core/logger"
	core_http_request "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/request"
	core_http_response "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/response"
)

type GetTaskResponse TaskDTOResponse

func (h *TaskHTTPHandler) GetTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskId, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task_id path value")
		return
	}

	task, err := h.taskService.GetTask(ctx, taskId)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task")
		return
	}

	response := GetTaskResponse(taskDTOFromDomain(task))

	responseHandler.JSONResponse(response, http.StatusOK)
}
