package middleware

import (
	"context"

	"../../database"
	"../../dcontext"
	"../../network"
)

type middleware struct {
	userRepository database.UserRepository
}

type MiddleWare interface {
	UserAuthorize(ar network.ApiResponser) network.ApiResponser
}

func NewMiddleWare(db database.ConnectedDB) MiddleWare {
	return &middleware{
		userRepository: database.NewUserRepository(db),
	}
}

func (mw *middleware) UserAuthorize(ar network.ApiResponser) network.ApiResponser {
	ctx := ar.GetRequestContext()
	if ctx == nil {
		ctx = context.Background()
	}

	token := ar.GetRequest().GetHeaderValue("x-token")
	if len(token) == 0 {
		ar.BadRequest("x-token is empty")
	}

	user, err := mw.userRepository.FindByAuthToken(token)
	if err != nil {
		ar.InternalServerError("User is not found: Not matching token found")
	}

	ctx = dcontext.SetUserID(ctx, user.UserID)

	ar.SetRequestContext(ctx)

	return ar
}
