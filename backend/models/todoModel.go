package models

import (
	"gorm.io/gorm"
)

type Models interface {
	GetSubject() string
}

type Todo struct {
	gorm.Model
	Subject string `json:"subject"`
}

func (m *Todo) GetSubject() string {
	return m.Subject
}
