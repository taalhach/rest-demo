package models

import "time"

type User struct {
	Id            int64     `json:"id" gorm:"primaryKey"`
	Uuid          string    `xorm:"uuid"`
	Role          string    `xorm:"role"`
	Email         string    `xorm:"email"`
	Password      string    `xorm:"password"`
	CreatedAt     time.Time `xorm:"created_at"`
	LastUpdatedAt time.Time `xorm:"last_updated_at"`
}

func (this User) TableName() string {
	return "users"
}
