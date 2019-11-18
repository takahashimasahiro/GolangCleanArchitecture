package router

import (
	"../server"
	"../../interface/controllers"
)

func BootRouter(s server.Server, controller controllers.AppController) {
	// auth
	s.Post("/auth/create", func(hc *server.HttpContext) { controller.CreateUser(hc) })
	// user
	s.Get("/user/get", func(hc *server.HttpContext) { controller.GetUser(hc) })
	s.Post("/user/update", func(hc *server.HttpContext) { controller.UpdateUser(hc) })
}