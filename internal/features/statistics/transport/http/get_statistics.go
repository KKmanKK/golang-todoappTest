package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/KKmanKK/golang-todoappTest/internal/core/domain"
	core_logger "github.com/KKmanKK/golang-todoappTest/internal/core/logger"
	core_http_request "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/request"
	core_http_response "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/response"
)

type GetStatisticsResponse struct {
	TaskCreated               int      `json:"task_created"`
	TaskCompleted             int      `json:"task_completed"`
	TaskCompletedRate         *float64 `json:"task_completed_rate"`
	TaskAverageCompletionTime *string  `json:"task_average_completion_time"`
}

func (h *StatisticsHTTPHandler) GetStatistics(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userId, from, to, err := getUserIdFromToQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user_id/from/to query param",
		)
		return
	}

	statisics, err := h.statisticSerivce.GetStatistics(
		ctx,
		userId,
		from,
		to,
	)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed get statistics")
		return
	}

	response := statisticsDTOFromDomain(statisics)

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getUserIdFromToQueryParams(r *http.Request) (*int, *time.Time, *time.Time, error) {
	const (
		userIdQueryParamKey = "user_id"
		FromQueryParamKey   = "from"
		ToQueryParamKey     = "to"
	)

	userId, err := core_http_request.GetIntQuertyParam(r, userIdQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}

	from, err := core_http_request.GetDateQueryParam(r, FromQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'from' query param: %w", err)
	}

	to, err := core_http_request.GetDateQueryParam(r, ToQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'to' query param: %w", err)
	}

	return userId, from, to, nil
}

func statisticsDTOFromDomain(statisics domain.Statistics) GetStatisticsResponse {
	var avgTime *string
	if statisics.TaskAverageCompletionTime != nil {
		duration := statisics.TaskAverageCompletionTime.String()
		avgTime = &duration
	}

	return GetStatisticsResponse{
		TaskCreated:               statisics.TaskCreated,
		TaskCompleted:             statisics.TaskCompleted,
		TaskCompletedRate:         statisics.TaskCompletedRate,
		TaskAverageCompletionTime: avgTime,
	}
}
