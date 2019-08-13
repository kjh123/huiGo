package cmd

import (
	"github.com/gorilla/mux"
	"github.com/kjh123/huiGo/log"
	"github.com/kjh123/huiGo/services"
	"github.com/urfave/negroni"
	"gopkg.in/tylerb/graceful.v1"
	"strconv"
	"time"
)

func Serve(configFile string) error {
	conf, db, err := initConfigDB(configFile)
	if err != nil {
		return err
	}
	defer db.Close()

	// start the services
	if err := services.Init(conf, db); err != nil {
		return err
	}

	// new negroni
	app := negroni.New()
	app.Use(negroni.NewRecovery())
	app.Use(negroni.NewLogger())

	router := mux.NewRouter()
	for _, s := range services.Modules() {
		s.GetRoutes()
		s.RegisterRoutes(router, s.Prefix())
	}
	defer services.Close()

	// Set the router
	app.UseHandler(router)

	log.INFO.Println("Starting server on port ", conf.ServerPort)
	// Run the server on $ServerPort, gracefully stop on SIGTERM signal
	graceful.Run(":"+strconv.Itoa(conf.ServerPort), 5*time.Second, app)

	return nil
}
