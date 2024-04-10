package model

import "gorm.io/gorm"

type Quiz struct {
  gorm.Model
  Name string `gorm:"not null"`
  Description string `gorm:"not null"`
}
