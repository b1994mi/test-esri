package complaint

import (
	"net/http"

	"github.com/b1994mi/test-esri/internal/pkg/usecase/complaint"
	"github.com/uptrace/bunrouter"
)

type Handler interface {
	CreateHandler(w http.ResponseWriter, bunReq bunrouter.Request) error
	FindPagedHandler(w http.ResponseWriter, bunReq bunrouter.Request) error
}

type handler struct {
	uc complaint.Usecase
}

func NewHandler(uc complaint.Usecase) *handler {
	return &handler{uc: uc}
}
