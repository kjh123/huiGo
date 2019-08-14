package user

import (
    "github.com/gorilla/mux"
    "github.com/kjh123/huiGo/util/routes"
    "github.com/urfave/negroni"
)

func (s *Service) RegisterRoutes(router *mux.Router, prefix string) {
    subRouter := router.PathPrefix(prefix).Subrouter()
    routes.AddRoutes(s.GetRoutes(), subRouter)
}

func (s *Service) GetRoutes() []routes.Route {
    return []routes.Route{
        {
            Name:        "get_users",
            Method:      "GET",
            Pattern:     "/users",
            HandlerFunc: s.getUsers,
        },
        {
            Name:        "create_user",
            Method:      "POST",
            Pattern:     "/users",
            HandlerFunc: s.createUser,
            Middlewares: []negroni.Handler{
                new(parseFormMiddleWare),
            },
        },
        {
            Name:        "update_user",
            Method:      "PUT",
            Pattern:     "/users",
            HandlerFunc: s.updateUser,
            Middlewares: []negroni.Handler{
                new(parseFormMiddleWare),
            },
        },
        {
            Name:        "delete_user",
            Method:      "DELETE",
            Pattern:     "/users",
            HandlerFunc: s.deleteUser,
            Middlewares: []negroni.Handler{
                new(parseFormMiddleWare),
            },
        },
    }
}
