package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/b1994mi/test-esri/internal/pkg/domain/helper"
	"github.com/uptrace/bunrouter"
)

func AuthMiddleware(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		token := req.Header["Authorization"]
		if len(token) < 1 {
			return next(w, req)
		}

		splitToken := strings.Split(token[0], " ")
		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			return next(w, req)
		}

		userID, err := helper.ParseToken(splitToken[1])
		if err != nil {
			log.Println(err)
			return next(w, req)
		}

		ctx := req.Context()
		ctx = context.WithValue(ctx, "user_id", userID)
		return next(w, req.WithContext(ctx))
	}
}
