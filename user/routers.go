package user

import (
	"github.com/gorilla/mux"
	"github.com/kjh123/huiGo/util/routes"
)

func (s *Service) RegisterRoutes(router *mux.Router, prefix string) {
	subRouter := router.PathPrefix(prefix).Subrouter()
	routes.AddRoutes(s.GetRoutes(), subRouter)
}

func (s *Service) GetRoutes() []routes.Route {
	return []routes.Route{
		{
			Name: "create_user",
			Method: "POST",
			Pattern: "/users",
			HandlerFunc: s.createUser,
		},
		{
			Name: "update_user",
			Method: "PUT",
			Pattern: "/user/{id}",
			HandlerFunc: s.updateUser,
		},
	}
}
