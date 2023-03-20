package route

import (
	"github.com/b1994mi/test-esri/internal/pkg/usecase/auth"
	"github.com/b1994mi/test-esri/internal/pkg/usecase/complaint"
)

type usecaseContainer struct {
	authUsecase      auth.Usecase
	complaintUsecase complaint.Usecase
}

func SetupUsecase(svc sqlserviceContainer) usecaseContainer {
	authUsecase := auth.NewUsecase(
		svc.UserService,
	)

	complaintUsecase := complaint.NewUsecase(
		svc.IssueService,
		svc.IssueImageService,
	)

	return usecaseContainer{
		authUsecase,
		complaintUsecase,
	}
}
