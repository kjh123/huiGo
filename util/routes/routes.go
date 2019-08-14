package routes

import (
    "github.com/gorilla/mux"
    "github.com/urfave/negroni"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
    Middlewares []negroni.Handler
}

func AddRoutes(routes []Route, router *mux.Router) {
    var (
        handler http.Handler
        n       *negroni.Negroni
    )

    for _, route := range routes {
        if len(route.Middlewares) > 0 {
            n = negroni.New()

            for _, middleware := range route.Middlewares {
                n.Use(middleware)
            }
            n.Use(negroni.Wrap(route.HandlerFunc))
            handler = n
        } else {
            handler = route.HandlerFunc
        }

        router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
    }
}
