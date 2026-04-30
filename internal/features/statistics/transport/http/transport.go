package statistics_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/KKmanKK/golang-todoappTest/internal/core/domain"
	core_http_server "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/server"
)

type StatisticsHTTPHandler struct {
	statisticSerivce StatisticSerivce
}

type StatisticSerivce interface {
	GetStatistics(
		ctx context.Context,
		userId *int,
		from *time.Time,
		to *time.Time,
	) (domain.Statistics, error)
}

func NewStatisticsHTTPHandler(
	statisticSerivce StatisticSerivce,
) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{
		statisticSerivce: statisticSerivce,
	}
}

func (h *StatisticsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistics",
			Handler: h.GetStatistics,
		},
	}
}
