package models

import (
	"gorm.io/gorm"
)

type Done struct {
	gorm.Model
	Subject string `json:"subject"`
}

func (m *Done) GetSubject() string {
	return m.Subject
}
