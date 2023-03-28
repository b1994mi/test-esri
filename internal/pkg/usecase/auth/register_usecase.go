package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/b1994mi/test-esri/internal/pkg/domain/sqlmodel"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"passsword"`
}

func (uc *usecase) RegisterUsecase(req RegisterRequest) (interface{}, error) {
	user, err := uc.userService.FindOneBy(map[string]interface{}{
		"username": req.Username,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("unable to find user: %v", err)
	}

	if user != nil {
		return nil, fmt.Errorf("username has been used")
	}

	b := []byte(req.Password)
	password, err := bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("unable to hash password: %v", err)
	}

	tx := uc.userService.StartTx()
	defer tx.Rollback()

	userRes, err := uc.userService.Create(&sqlmodel.User{
		FullName:       req.FullName,
		Username:       req.Username,
		Email:          req.Email,
		PasswordHash:   string(password),
		AvatarFileName: "",
		IsDeleted:      false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, tx)
	if err != nil {
		return nil, fmt.Errorf("unable to create user: %v", err)
	}

	tx.Commit()

	return userRes, nil
}
