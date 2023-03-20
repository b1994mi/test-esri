package complaint

import (
	"net/http"
	"strconv"

	"github.com/b1994mi/test-esri/internal/pkg/usecase/complaint"
	"github.com/uptrace/bunrouter"
)

func (h *handler) FindPagedHandler(w http.ResponseWriter, bunReq bunrouter.Request) error {
	userID := bunReq.Context().Value("user_id")
	if userID == nil {
		w.WriteHeader(http.StatusUnauthorized)
		bunrouter.JSON(w, bunrouter.H{
			"code":    "007",
			"message": "",
		})

		return nil
	}

	pageStr := bunReq.URL.Query().Get("page")
	page, _ := strconv.Atoi(pageStr)
	sizeStr := bunReq.URL.Query().Get("size")
	size, _ := strconv.Atoi(sizeStr)

	req := complaint.PagedComplaintRequest{
		// Query: paramMap["query"],
		Page:              page,
		Size:              size,
		AuthenticatedUser: userID.(int),
	}

	res, err := h.uc.FindPagedUsecase(req)
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
