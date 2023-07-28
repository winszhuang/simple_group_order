package models

import "time"

type User struct {
	UserID    int       `gorm:"primary_key;auto_increment"`
	Username  string    `gorm:"varchar"`
	Password  string    `gorm:"varchar"`
	Email     string    `gorm:"varchar"`
	CreatedAt time.Time `gorm:"datetime"`
}
