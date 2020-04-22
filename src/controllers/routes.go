package controllers

import "woden/src/middlewares"

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/user", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/user/auth", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/user", middlewares.SetMiddlewareJSON(s.UpdateUser)).Methods("PUT")
	s.Router.HandleFunc("/user/logout", middlewares.SetMiddlewareJSON(s.Logout)).Methods("DELETE")
}
