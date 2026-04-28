package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_logger "github.com/KKmanKK/golang-todoappTest/internal/core/logger"
	core_pgx_pool "github.com/KKmanKK/golang-todoappTest/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/middleware"
	core_http_server "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/server"
	task_postgres_repository "github.com/KKmanKK/golang-todoappTest/internal/features/task/repository/postgres"
	task_service "github.com/KKmanKK/golang-todoappTest/internal/features/task/service"
	task_transport_http "github.com/KKmanKK/golang-todoappTest/internal/features/task/transport/http"
	users_postgres_repository "github.com/KKmanKK/golang-todoappTest/internal/features/users/repository/postgres"
	users_service "github.com/KKmanKK/golang-todoappTest/internal/features/users/service"
	user_transport_http "github.com/KKmanKK/golang-todoappTest/internal/features/users/transport/http"
	"go.uber.org/zap"
)

var (
	timeZone = time.UTC
)

func main() {
	time.Local = timeZone

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())

	if err != nil {
		fmt.Println("faliled to init appl;ication logger: ", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("application time zone", zap.Any("zone", timeZone))

	logger.Debug("initializing postgres connection pool")
	pool, err := core_pgx_pool.NewPool(
		ctx,
		core_pgx_pool.NewConfigMust(),
	)

	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("inintializing feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUserRepository(pool)
	userService := users_service.NewUsersSevice(usersRepository)
	usersTransportHTTP := user_transport_http.NewUsersHTTPHander(userService)

	logger.Debug("inintializing feature", zap.String("feature", "tasks"))
	taskRepository := task_postgres_repository.NewTaskRepository(pool)
	taskService := task_service.NewTaskService(taskRepository)
	taskTransportHTTP := task_transport_http.NewTaskHTTPHander(taskService)

	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)
	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRouters(usersTransportHTTP.Routes()...)
	apiVersionRouter.RegisterRouters(taskTransportHTTP.Routes()...)
	httpServer.RegisterApiRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
