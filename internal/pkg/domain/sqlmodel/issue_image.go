package sqlmodel

import (
	"time"
)

type IssueImage struct {
	ID        int       `json:"id" gorm:"primarykey"`
	IssueID   int       `json:"issue_id"`
	Filename  string    `json:"filename"`
	IsPrimary bool      `json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *IssueImage) TableName() string {
	return "issue_images"
}
