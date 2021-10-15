package models

import (
	"time"
)

type Task struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	AssignedTo string    `json:"assignedTo"`
	Task       string    `json:"task"`
	Deadline   time.Time `json:"deadline"`
	CreateAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"update_at"`
}
