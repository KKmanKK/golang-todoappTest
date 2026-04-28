package task_transport_http

import (
	"fmt"
	"net/http"

	"github.com/KKmanKK/golang-todoappTest/internal/core/domain"
	core_logger "github.com/KKmanKK/golang-todoappTest/internal/core/logger"
	core_http_request "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/request"
	core_http_response "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/response"
	core_http_types "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/types"
)

type PatchTaskRequest struct {
	Title       core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	Completed   core_http_types.Nullable[bool]   `json:"completed"`
}

func (r *PatchTaskRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("'Title' can't be null")
		}

		titleLen := len([]rune(*r.Title.Value))
		if titleLen < 1 || titleLen > 100 {
			return fmt.Errorf("'Title' must be between 1 and 100 symbols")
		}
	}

	if r.Description.Set {
		if r.Description.Value != nil {
			descriptionLen := len([]rune(*r.Description.Value))
			if descriptionLen < 1 || descriptionLen > 1000 {
				return fmt.Errorf("'Description' must be between 1 and 1000 symbols")
			}
		}
	}

	if r.Completed.Set {
		if r.Completed.Value == nil {
			return fmt.Errorf("'Completed' can't be null")
		}
	}

	return nil
}

type PatchTaskResponse TaskDTOResponse

func (h *TaskHTTPHandler) PatchTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskId, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task_id path value")
		return
	}

	var request PatchTaskRequest
	err = core_http_request.DecodeAndValidateRequest(r, &request)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	taskPatch := taskPatchFromRequest(request)

	task, err := h.taskService.PatchTask(ctx, taskId, taskPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch task")
		return
	}

	response := PatchTaskResponse(taskDTOFromDomain(task))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func taskPatchFromRequest(request PatchTaskRequest) domain.TaskPatch {
	return domain.NewTaskPatch(
		request.Title.ToDomain(),
		request.Description.ToDomain(),
		request.Completed.ToDomain(),
	)
}
