package user

import (
	"github.com/gorilla/mux"
	"github.com/kjh123/huiGo/util/routes"
)

type ServiceInterface interface {
	GetRoutes() []routes.Route
	RegisterRoutes(router *mux.Router, prefix string)
	Close()
}
