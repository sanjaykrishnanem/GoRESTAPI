package Models

type User struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Rollnum string `json:"rollnum" gorm:"primary_key;unique"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   uint   `json:"phone"`
	Address string `json:"address"`
}

func (b *User) TableName() string {
	return "user"
}
