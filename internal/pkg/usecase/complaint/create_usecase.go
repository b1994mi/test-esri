package complaint

import (
	"fmt"
	"time"

	"github.com/b1994mi/test-esri/internal/pkg/domain/sqlmodel"
)

type CreateComplaintRequest struct {
	MeteranID         int    `json:"meteran_id"`
	CategoryID        int    `json:"category_id"`
	ComplaintName     string `json:"complaint_name"`
	ShortDescription  string `json:"short_description"`
	PriorityLevel     int    `json:"priority_level"`
	Media             []string
	AuthenticatedUser int
}

func (uc *usecase) CreateUsecase(req CreateComplaintRequest) (interface{}, error) {
	tx := uc.issueService.StartTx()
	defer tx.Rollback()

	issue, err := uc.issueService.Create(&sqlmodel.Issue{
		UserID:           req.AuthenticatedUser,
		MeteranID:        req.MeteranID,
		CategoryID:       req.CategoryID,
		ComplaintName:    req.ComplaintName,
		ShortDescription: req.ShortDescription,
		PriorityLevel:    req.PriorityLevel,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}, tx)
	if err != nil {
		return nil, fmt.Errorf("unable to create issue: %v", err)
	}

	for i, media := range req.Media {
		_, err := uc.issueImageService.Create(&sqlmodel.IssueImage{
			IssueID:   issue.ID,
			Filename:  media,
			IsPrimary: i == 0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, tx)
		if err != nil {
			return nil, fmt.Errorf("unable to create issue image: %v", err)
		}
	}

	tx.Commit()

	return issue, nil
}
