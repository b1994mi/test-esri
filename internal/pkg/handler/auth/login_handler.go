package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/b1994mi/test-esri/internal/pkg/usecase/auth"
	"github.com/uptrace/bunrouter"
)

func (h *handler) LoginHandler(w http.ResponseWriter, bunReq bunrouter.Request) error {
	body, err := ioutil.ReadAll(bunReq.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		bunrouter.JSON(w, bunrouter.H{
			"code":    "007",
			"message": err,
		})

		return nil
	}

	var req auth.LoginRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		bunrouter.JSON(w, bunrouter.H{
			"code":    "007",
			"message": err,
		})

		return nil
	}

	res, err := h.uc.LoginUsecase(req)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		bunrouter.JSON(w, bunrouter.H{
			"code":    "007",
			"message": err,
		})

		return nil
	}

	bunrouter.JSON(w, bunrouter.H{
		"data": res,
	})
	return nil
}
