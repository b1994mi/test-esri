package sqlmodel

import (
	"time"
)

type Issue struct {
	ID               int       `json:"id" gorm:"primarykey"`
	UserID           int       `json:"user_id"`
	MeteranID        int       `json:"meteran_id"`
	CategoryID       int       `json:"category_id"`
	ComplaintName    string    `json:"complaint_name"`
	ShortDescription string    `json:"short_description"`
	PriorityLevel    int       `json:"priority_level"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (m *Issue) TableName() string {
	return "issues"
}

type IssueWithImageMeteranCategory struct {
	Issue
	Images   []*IssueImage  `json:"images" gorm:"foreignKey:IssueID"`
	Meteran  *Meteran       `json:"meteran" gorm:"foreignKey:MeteranID"`
	Category *IssueCategory `json:"category" gorm:"foreignKey:CategoryID"`
}
