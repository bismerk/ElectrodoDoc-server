package controllers

import (
	"woden/src/middlewares"
)

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/api/v1/user", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/api/v1/user/auth", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/api/v1/user", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/api/v1/user/logout", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.Logout))).Methods("DELETE")

}
