package user_transport_http

import (
	"context"
	"net/http"

	"github.com/KKmanKK/golang-todoappTest/internal/core/domain"
	core_http_server "github.com/KKmanKK/golang-todoappTest/internal/core/transport/http/server"
)

type UsersHTTPHander struct {
	usersService UserService
}

type UserService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)
	GetUser(
		ctx context.Context,
		id int,
	) (domain.User, error)
	DeleteUser(
		ctx context.Context,
		id int,
	) error
	PatchUser(
		ctx context.Context,
		id int,
		user domain.UserPatch,
	) (domain.User, error)
}

func NewUsersHTTPHander(userService UserService) *UsersHTTPHander {
	return &UsersHTTPHander{
		usersService: userService,
	}
}

func (h *UsersHTTPHander) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUsers,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.DeleteUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: h.PatchUser,
		},
	}
}
