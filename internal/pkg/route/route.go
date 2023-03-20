package route

import (
	"net/http"

	"github.com/b1994mi/test-esri/internal/pkg/middleware"
	"github.com/uptrace/bunrouter"
)

type Router struct {
	bunrouter.Router
}

func NewRouter(h handlerContainer) *bunrouter.Router {
	r := bunrouter.New()

	r.GET("/", func(w http.ResponseWriter, bunReq bunrouter.Request) error {
		bunrouter.JSON(w, bunrouter.H{
			"message": "pong",
		})
		return nil
	})

	r.POST("/register", h.authHandler.RegisterHandler)
	r.POST("/login", h.authHandler.LoginHandler)
	r.Use(middleware.AuthMiddleware).GET("/complaint", h.complaintHandler.FindPagedHandler)
	r.Use(middleware.AuthMiddleware).POST("/complaint", h.complaintHandler.CreateHandler)

	return r
}
