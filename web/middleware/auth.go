package middleware

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

type Authenticate struct {
}

func NewAuthenticator() *Authenticate {
	return &Authenticate{}
}

func (a *Authenticate) Middleware() negroni.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		log.Printf("do some authentication here")
		next(rw, r)
	}
}
