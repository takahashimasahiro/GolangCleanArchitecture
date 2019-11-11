package controllers

import (
	"../controllers/middleware"
	"../database"
	"../network"
)

type interactor struct {
	db database.ConnectedDB
}

type Interactor interface {
	NewAppController() AppController
}

func NewInteractor(db database.ConnectedDB) Interactor {
	return &interactor{db: db}
}

func (i *interactor) NewAppController() AppController {
	return &appController{
		middleware:     middleware.NewMiddleWare(i.db),
		authController: NewAuthController(i.db),
		userController: NewUserController(i.db),
	}
}

type appController struct {
	middleware     middleware.Middleware
	authController AuthController
	userController UserController
}

type AppController interface {
	// authController
	CreateUser(ar network.ApiResponser)
	// userController
	GetUser(ar network.ApiResponser)
	UpdateUser(ar network.ApiResponser)
}

func (ac *appController) CreateUser(ar network.ApiResponser) {

}

func (ac *appController) GetUser(ar network.ApiResponser) {

}

func (ac *appController) UpdateUser(ar network.ApiResponser) {

}
