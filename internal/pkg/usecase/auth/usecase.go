package auth

import (
	"github.com/b1994mi/test-esri/internal/pkg/domain/sqlservice"
)

type Usecase interface {
	LoginUsecase(req LoginRequest) (interface{}, error)
}

type usecase struct {
	userService sqlservice.UserService
}

func NewUsecase(
	userService sqlservice.UserService,
) *usecase {
	return &usecase{
		userService,
	}
}
