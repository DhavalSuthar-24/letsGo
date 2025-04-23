package models

import "gorm.io/gorm"

type Todo struct{
	gorm.Model
	Title string  `json:"title"`
	Description string `json:"description"`
	isCompleted bool `json:"completed"`
	UserID uint `json:"user_id"` // ✅ exported and usable by GORM

}