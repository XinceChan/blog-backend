package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"Description"`
	Image       string `json:"image"`
	UserID      string `json:"userid"`
	User        User   `json:"user" gorm:"foreignKey:UserID"`
}
