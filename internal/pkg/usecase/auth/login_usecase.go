package auth

import (
	"fmt"

	"github.com/b1994mi/test-esri/internal/pkg/domain/helper"
)

type LoginRequest struct {
	Username string
	Password string
}

func (uc *usecase) LoginUsecase(req LoginRequest) (interface{}, error) {
	user, err := uc.userService.FindOneBy(map[string]interface{}{
		"username": req.Username,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to find user: %v", err)
	}

	if user.IsDeleted {
		return nil, fmt.Errorf("this account has been deleted")
	}

	if !helper.IsCorrectPass(user.PasswordHash, req.Password) {
		return nil, fmt.Errorf("wrong password")
	}

	return helper.GenerateToken(user.ID), nil
}
