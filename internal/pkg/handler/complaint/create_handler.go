package complaint

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/b1994mi/test-esri/internal/pkg/domain/helper"
	"github.com/b1994mi/test-esri/internal/pkg/usecase/complaint"
	"github.com/uptrace/bunrouter"
)

// const maxFileSize = 3 * 1024 * 1024

func (h *handler) CreateHandler(w http.ResponseWriter, bunReq bunrouter.Request) error {
	err := bunReq.ParseMultipartForm(64)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		bunrouter.JSON(w, bunrouter.H{
			"code":    "007",
			"message": err,
		})

		return nil
	}

	var req complaint.CreateComplaintRequest
	for _, v := range bunReq.MultipartForm.Value["issue"] {
		json.Unmarshal([]byte(v), &req)
		break
	}

	req.AuthenticatedUser = bunReq.Context().Value("user_id").(int)

	wd, err := os.Getwd()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		bunrouter.JSON(w, bunrouter.H{
			"code":    "007",
			"message": err,
		})

		return nil
	}

	files := bunReq.MultipartForm.File["image"]
	for _, f := range files {
		dir := path.Join(wd, "/tmp/media")
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				bunrouter.JSON(w, bunrouter.H{
					"code":    "007",
					"message": err,
				})

				return nil
			}
		}

		tmpFile := path.Join(dir, fmt.Sprintf("%v-%v", time.Now().Unix(), f.Filename))
		if err = helper.SaveUploadedFile(f, tmpFile); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			bunrouter.JSON(w, bunrouter.H{
				"code":    "007",
				"message": err,
			})

			return nil
		}

		req.Media = append(req.Media, tmpFile)
	}

	res, err := h.uc.CreateUsecase(req)
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
