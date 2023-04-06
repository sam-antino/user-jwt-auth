package entities

import "time"

type Users struct {
	ID        int64  `gorm:"column:id;primaryKey;auto_increment;not null"`
	Name      string `gorm:"column:name;not null"`
	Email     string `gorm:"column:email;not null;unique"`
	Password  string `gorm:"column:password;not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
