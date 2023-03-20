package sqlservice

import (
	"github.com/b1994mi/test-esri/internal/pkg/domain/helper"
	"github.com/b1994mi/test-esri/internal/pkg/domain/sqlmodel"
	"gorm.io/gorm"
)

type UserService interface {
	StartTx() *gorm.DB
	Create(m *sqlmodel.User, tx *gorm.DB) (*sqlmodel.User, error)
	Update(m *sqlmodel.User, tx *gorm.DB) error
	Delete(m *sqlmodel.User, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*sqlmodel.User, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*sqlmodel.User, error)
	Count(criteria map[string]interface{}) int64
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *userService {
	return &userService{
		db,
	}
}

func (svc *userService) StartTx() *gorm.DB {
	return svc.db.Begin()
}

func (svc *userService) Create(m *sqlmodel.User, tx *gorm.DB) (*sqlmodel.User, error) {
	err := tx.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (svc *userService) Update(m *sqlmodel.User, tx *gorm.DB) error {
	return tx.Save(&m).Error
}

func (svc *userService) Delete(m *sqlmodel.User, tx *gorm.DB) error {
	return tx.Delete(&m).Error
}

func (svc *userService) FindOneBy(criteria map[string]interface{}) (*sqlmodel.User, error) {
	var m sqlmodel.User

	err := svc.db.Where(criteria).Take(&m).Error
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (svc *userService) FindBy(criteria map[string]interface{}, page, size int) ([]*sqlmodel.User, error) {
	var data []*sqlmodel.User
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

func (svc *userService) Count(criteria map[string]interface{}) int64 {
	var result int64

	if res := svc.db.Model(sqlmodel.User{}).Where(criteria).Count(&result); res.Error != nil {
		return 0
	}

	return result
}
