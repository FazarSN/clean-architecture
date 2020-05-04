package handler

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler func(w http.ResponseWriter, r *http.Request, params httprouter.Params) error

func Middleware(handle Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		log.Println(r.URL)
		handle(w, r, params)
	}
}
