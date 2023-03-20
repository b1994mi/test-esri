package route

import (
	"github.com/b1994mi/test-esri/internal/pkg/handler/auth"
	"github.com/b1994mi/test-esri/internal/pkg/handler/complaint"
)

type handlerContainer struct {
	authHandler      auth.Handler
	complaintHandler complaint.Handler
}

func SetupHandler(uc usecaseContainer) handlerContainer {
	return handlerContainer{
		auth.NewHandler(uc.authUsecase),
		complaint.NewHandler(uc.complaintUsecase),
	}
}
