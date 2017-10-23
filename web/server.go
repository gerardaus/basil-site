package web

import (
	"fmt"
	"net/http"

	"github.com/gerardaus/basil-site/controllers"
	"github.com/gerardaus/basil-site/web/middleware"
	oauth2 "github.com/goincremental/negroni-oauth2"
	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Server struct {
	*negroni.Negroni
}

func NewServer() *Server {
	s := Server{negroni.Classic()}
	router := mux.NewRouter()

	// https://github.com/xyproto/permissions2
	// New permissions middleware
	//perm, err := permissions.New2()
	//if err != nil {
	//	log.Fatalln(err)
	//}

	secureMux := http.NewServeMux()

	// Routes that require a logged in user
	// can be protected by using a separate route handler
	// If the user is not authenticated, they will be
	// redirected to the login path.
	secureMux.HandleFunc("/restrict", func(w http.ResponseWriter, req *http.Request) {
		token := oauth2.GetToken(req)
		fmt.Fprintf(w, "OK: %s", token.Access())
	})

	secure := negroni.New()
	secure.Use(oauth2.LoginRequired())
	secure.UseHandler(secureMux)

	router.Handle("/restrict", secure)

	// controllers
	homeController := controllers.NewHomeController()
	homeController.Register(router)

	s.Use(sessions.Sessions("the_session3", cookiestore.New([]byte("secret123"))))

	s.Use(middleware.NewAuthenticator().Middleware())
	s.Use(oauth2.Facebook(&oauth2.Config{
		ClientID:     "356566358130918",
		ClientSecret: "5c5c52c60c27a9782e2eed04ba9e3f09",
		RedirectURL:  "http://localhost:3000",
		Scopes:       []string{"email"},
	}))
	s.UseHandler(router)
	return &s
}
