package controllers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type HomeControllerImpl struct {
}

func NewHomeController() *HomeControllerImpl {
	return &HomeControllerImpl{}
}

func (hc *HomeControllerImpl) Register(router *mux.Router) {
	router.HandleFunc("/", hc.single)
}

func (hc *HomeControllerImpl) single(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("views/home.html")
	t.Execute(w, nil)
}
