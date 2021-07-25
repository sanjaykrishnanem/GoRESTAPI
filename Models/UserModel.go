package Models

import "time"

type User struct {
	Email        string `json:"email;" gorm:"primary_key" sql:"not null;unique"`
	Name         string `json:"name;" sql:"not null;"`
	Phone        uint64 `json:"phone;" sql:"not null;unique;"`
	Verified     bool   `json:"verified;" sql:"DEFAULT:false;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	PasswordHash string `json:"password"`
}

func (b *User) TableName() string {
	return "user"
}
