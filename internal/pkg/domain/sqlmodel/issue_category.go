package sqlmodel

import (
	"time"
)

type IssueCategory struct {
	ID        int       `json:"id" gorm:"primarykey"`
	Name      string    `json:"name"`
	Class     string    `json:"class"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *IssueCategory) TableName() string {
	return "issue_categories"
}
