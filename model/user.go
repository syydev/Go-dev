package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserID   string `gorm:"primary_key;unique_index;not null" json:"userId"`
	Username string `gorm:"not null" json:"userName"`
	Password string `gorm:"not null" json:"password"`
}

func (u *User) Create() error {
	return db.local.Create(&u).Error
}
