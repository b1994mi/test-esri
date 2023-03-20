package complaint

import "fmt"

type PagedComplaintRequest struct {
	Query             string
	Size              int
	Page              int
	AuthenticatedUser int
}

func (uc *usecase) FindPagedUsecase(req PagedComplaintRequest) (interface{}, error) {
	issues, err := uc.issueService.FindWithImageMeteranCategoryBy(map[string]interface{}{}, req.Page, req.Size)
	if err != nil {
		return nil, fmt.Errorf("unable to find complaint list: %v", err)
	}

	return issues, nil
}
