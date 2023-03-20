package sqlservice

import (
	"github.com/b1994mi/test-esri/internal/pkg/domain/helper"
	"github.com/b1994mi/test-esri/internal/pkg/domain/sqlmodel"
	"gorm.io/gorm"
)

type IssueService interface {
	StartTx() *gorm.DB
	Create(m *sqlmodel.Issue, tx *gorm.DB) (*sqlmodel.Issue, error)
	Update(m *sqlmodel.Issue, tx *gorm.DB) error
	Delete(m *sqlmodel.Issue, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*sqlmodel.Issue, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*sqlmodel.Issue, error)
	Count(criteria map[string]interface{}) int64
	FindWithImageMeteranCategoryBy(criteria map[string]interface{}, page, size int) ([]*sqlmodel.IssueWithImageMeteranCategory, error)
}

type issueService struct {
	db *gorm.DB
}

func NewIssueService(db *gorm.DB) *issueService {
	return &issueService{
		db,
	}
}

func (svc *issueService) StartTx() *gorm.DB {
	return svc.db.Begin()
}

func (svc *issueService) Create(m *sqlmodel.Issue, tx *gorm.DB) (*sqlmodel.Issue, error) {
	err := tx.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (svc *issueService) Update(m *sqlmodel.Issue, tx *gorm.DB) error {
	return tx.Save(&m).Error
}

func (svc *issueService) Delete(m *sqlmodel.Issue, tx *gorm.DB) error {
	return tx.Delete(&m).Error
}

func (svc *issueService) FindOneBy(criteria map[string]interface{}) (*sqlmodel.Issue, error) {
	var m sqlmodel.Issue

	err := svc.db.Where(criteria).Take(&m).Error
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (svc *issueService) FindBy(criteria map[string]interface{}, page, size int) ([]*sqlmodel.Issue, error) {
	var data []*sqlmodel.Issue
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

func (svc *issueService) FindWithImageMeteranCategoryBy(criteria map[string]interface{}, page, size int) ([]*sqlmodel.IssueWithImageMeteranCategory, error) {
	var data []*sqlmodel.IssueWithImageMeteranCategory
	if page == 0 || size == 0 {
		page, size = -1, -1
	}

	limit, offset := helper.GetLimitOffset(page, size)
	err := svc.db.Debug().
		Where(criteria).
		Joins("Meteran").
		Joins("Category").
		Preload("Images").
		Offset(offset).Limit(limit).
		Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (svc *issueService) Count(criteria map[string]interface{}) int64 {
	var result int64

	if res := svc.db.Model(sqlmodel.Issue{}).Where(criteria).Count(&result); res.Error != nil {
		return 0
	}

	return result
}
