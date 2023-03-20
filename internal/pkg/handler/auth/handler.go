package auth

import (
	"net/http"

	"github.com/b1994mi/test-esri/internal/pkg/usecase/auth"
	"github.com/uptrace/bunrouter"
)

type Handler interface {
	LoginHandler(w http.ResponseWriter, bunReq bunrouter.Request) error
}

type handler struct {
	uc auth.Usecase
}

func NewHandler(uc auth.Usecase) *handler {
	return &handler{uc: uc}
}
