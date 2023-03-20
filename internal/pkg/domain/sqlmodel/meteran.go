package sqlmodel

import (
	"time"
)

type Meteran struct {
	ID          int       `json:"id" gorm:"primarykey"`
	MeteranCode string    `json:"meteran_code"`
	Address     string    `json:"address"`
	Lat         float64   `json:"lat"`
	Lon         float64   `json:"lon"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (m *Meteran) TableName() string {
	return "meterans"
}
