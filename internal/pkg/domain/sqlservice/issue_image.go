package sqlservice

import (
	"github.com/b1994mi/test-esri/internal/pkg/domain/helper"
	"github.com/b1994mi/test-esri/internal/pkg/domain/sqlmodel"
	"gorm.io/gorm"
)

type IssueImageService interface {
	StartTx() *gorm.DB
	Create(m *sqlmodel.IssueImage, tx *gorm.DB) (*sqlmodel.IssueImage, error)
	Update(m *sqlmodel.IssueImage, tx *gorm.DB) error
	Delete(m *sqlmodel.IssueImage, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*sqlmodel.IssueImage, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*sqlmodel.IssueImage, error)
	Count(criteria map[string]interface{}) int64
}

type issueImageService struct {
	db *gorm.DB
}

func NewIssueImageService(db *gorm.DB) *issueImageService {
	return &issueImageService{
		db,
	}
}

func (svc *issueImageService) StartTx() *gorm.DB {
	return svc.db.Begin()
}

func (svc *issueImageService) Create(m *sqlmodel.IssueImage, tx *gorm.DB) (*sqlmodel.IssueImage, error) {
	err := tx.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (svc *issueImageService) Update(m *sqlmodel.IssueImage, tx *gorm.DB) error {
	return tx.Save(&m).Error
}

func (svc *issueImageService) Delete(m *sqlmodel.IssueImage, tx *gorm.DB) error {
	return tx.Delete(&m).Error
}

func (svc *issueImageService) FindOneBy(criteria map[string]interface{}) (*sqlmodel.IssueImage, error) {
	var m sqlmodel.IssueImage

	err := svc.db.Where(criteria).Take(&m).Error
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (svc *issueImageService) FindBy(criteria map[string]interface{}, page, size int) ([]*sqlmodel.IssueImage, error) {
	var data []*sqlmodel.IssueImage
	if page == 0 || size == 0 {
		page, size = -1, -1
	}

	limit, offset := helper.GetLimitOffset(page, size)
	err := svc.db.
		Where(criteria).
		Offset(offset).Limit(limit).
		Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (svc *issueImageService) Count(criteria map[string]interface{}) int64 {
	var result int64

	if res := svc.db.Model(sqlmodel.IssueImage{}).Where(criteria).Count(&result); res.Error != nil {
		return 0
	}

	return result
}
