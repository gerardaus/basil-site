package web

import (
	"github.com/codegangsta/negroni"
	"github.com/gerardaus/basil-site/controllers"
	"github.com/gerardaus/basil-site/web/middleware"
	"github.com/gorilla/mux"
)

type Server struct {
	*negroni.Negroni
}

func NewServer() *Server {
	s := Server{negroni.Classic()}
	router := mux.NewRouter()

	homeController := controllers.NewHomeController()
	homeController.Register(router)

	s.Use(middleware.NewAuthenticator().Middleware())
	s.UseHandler(router)
	return &s
}
