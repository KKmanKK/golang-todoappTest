package task_transport_http

import (
	"net/http"

	"github.com/KKmanKK/golang-todoappTest/internal/core/domain"
	core_logger "github.com/KKmanKK/golang-todoappTest/internal/core/logger"
	core_http_request "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/request"
	core_http_response "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/response"
)

type CreateTaskRequest struct {
	Title        string  `json:"title" validate:"required,min=1,max=100"`
	Description  *string `json:"description" validate:"omitempty,min=1,max=1000"`
	AuthorUserId int     `json:"author_user_id" validate:"required"`
}

type CreateTaskResponse TaskDTOResponse

func (h *TaskHTTPHandler) CreateTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreateTaskRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	taskDomain := domainFromDTO(request)

	taskDomain, err := h.taskService.CreateTask(ctx, taskDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create task")
		return
	}

	response := CreateTaskResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateTaskRequest) domain.Task {
	return domain.NewTaskUninitialized(
		dto.Title,
		dto.Description,
		dto.AuthorUserId,
	)
}
