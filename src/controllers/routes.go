package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"woden/src/middlewares"
)

func (s *Server) initializeRoutes() {
	s.Router.PathPrefix("/api").Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "authorization,content-type")
		w.Header().Set("Access-Control-Expose-Headers", "content-disposition")
	})
	s.Router.Use(mux.CORSMethodMiddleware(s.Router))

	s.Router.HandleFunc("/api/v1/user", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/api/v1/user/auth", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/api/v1/user", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/api/v1/user/logout", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.Logout))).Methods("DELETE")

}
