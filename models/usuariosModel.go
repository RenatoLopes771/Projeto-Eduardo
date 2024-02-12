package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	ID    int //`gorm:"primaryKey;autoIncrement"`
	Nome  string
	Email string
	Senha string
}
