package controllers

import (
	"encoding/json"
	"log"

	"../../usecase/service"
	"../database"
	"../network"
)

type authController struct {
	authService service.AuthService
}

type AuthController interface {
	CreateUser(network.ApiResponser)
}

func NewAuthController(db database.ConnectedDB) AuthController {
	return &authController{
		authService: service.NewAuthService(
			database.NewUserRepository(db),
		),
	}
}

func (ac *authController) CreateUser(ar network.ApiResponser) {
	var authCreateRequest AuthCreateRequest
	err := json.NewDecoder(ar.GetRequest().GetBody()).Decode(&authCreateRequest)
	if err != nil {
		log.Println("%+v\n", err)
		ar.BadRequest("Invalid Request")
		return
	}
	authToken, err := ac.authService.CreateUser(&authCreateRequest.Name)
	if err != nil {
		return
	}

	authCreateResponse := AuthCreateResponse{
		Token: *authToken,
	}

	ar.Success(authCreateResponse)
}

type AuthCreateRequest struct {
	Name string `json:"name"`
}

type AuthCreateResponse struct {
	Token string `json:"token"`
}
