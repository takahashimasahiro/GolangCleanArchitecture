package service

import(
	"../../domain"
	"../repository"
	"log"

	"github.com/google/uuid"
)

type authService struct {
	UserRepository repository.UserRepository
}

type AuthService interface {
	CreateUser(userName *string) (*string, error)
}

func NewAuthService(ur repository.UserRepository) AuthService {
	return &authService{UserRepository: ur}
}

func (authService *authService) CreateUser(userName *string) (*string, error) {
	userID, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	userIDString := userID.String()

	authToken, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	authTokenString := authToken.String()

	user := domain.User{
		UserID: userIDString,
		AuthToken: authTokenString,
		Name: *userName,
	}
	err = authService.UserRepository.Store(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &authTokenString, nil
}