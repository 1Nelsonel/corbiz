package model

import (
    "gorm.io/gorm"
)

type Service struct {
    gorm.Model
    Name      string `gorm:"not null"`
    ImagePath string
    Content   string `gorm:"type:text"`
}