package route

import (
	"github.com/b1994mi/test-esri/internal/pkg/domain/sqlservice"
	"gorm.io/gorm"
)

type sqlserviceContainer struct {
	sqlservice.UserService
	sqlservice.IssueService
	sqlservice.IssueImageService
}

func SetupSQLService(db *gorm.DB) sqlserviceContainer {
	return sqlserviceContainer{
		sqlservice.NewUserService(db),
		sqlservice.NewIssueService(db),
		sqlservice.NewIssueImageService(db),
	}
}
