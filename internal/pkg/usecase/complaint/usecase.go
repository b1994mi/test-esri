package complaint

import (
	"github.com/b1994mi/test-esri/internal/pkg/domain/sqlservice"
)

type Usecase interface {
	CreateUsecase(req CreateComplaintRequest) (interface{}, error)
	FindPagedUsecase(req PagedComplaintRequest) (interface{}, error)
}

type usecase struct {
	issueService      sqlservice.IssueService
	issueImageService sqlservice.IssueImageService
}

func NewUsecase(
	issueService sqlservice.IssueService,
	issueImageService sqlservice.IssueImageService,
) *usecase {
	return &usecase{
		issueService,
		issueImageService,
	}
}
