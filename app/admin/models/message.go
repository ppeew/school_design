package models

import "go-admin/cmd/migrate/migration/models"

type Message struct {
	models.Model
	models.BaseModel
	UserID       int
	Type         int
	Content      string
	NoticeStatus int
}
