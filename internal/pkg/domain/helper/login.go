package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func IsCorrectPass(hashed string, plain string) bool {
	byteHash := []byte(hashed)
	bytePlain := []byte(plain)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		return false
	}

	return true
}

func GenerateToken(userID int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * (3)).Unix(),
	})

	tokenString, err := token.SignedString([]byte("h"))
	if err != nil {
		return ""
	}

	return tokenString
}

func ParseToken(tokenString string) (userID int, err error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return "", nil
		}

		return []byte("h"), nil
	})
	if err != nil || token == nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, err
	}

	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Unix() {
		return 0, err
	}

	userID = int(claims["user_id"].(float64))
	return userID, err
}
