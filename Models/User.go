package Models

import (
	"CRUDTEST/Config"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *User) (err error) {
	pass := user.PasswordHash
	if len(pass) == 0 {
		return errors.New("Password should not be empty!")
	}
	user.PasswordHash = ""
	if err = Config.DB.Create(user).Error; err != nil {
		fmt.Println(err)
		return err
	}
	if err = user.setPassword(pass); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *User, email string) (err error) {
	if err = Config.DB.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		fmt.Println("User not found")
		return err
	}
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}

func (u *User) setPassword(password string) error {
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	Config.DB.Save(u)
	return nil
}

// Database will only save the hashed string, you should check it by util function.
// 	if err := serModel.checkPassword("password0"); err != nil { password error }
func (u *User) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

//ValidateLogin : Login Validation
func ValidateLogin(email string, password string) (err error) {
	var user User
	if err = Config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		fmt.Println("User not found")
		return err
	}
	if err := user.checkPassword(password); err != nil {
		fmt.Println("password error")
		return err
	}
	return nil
}
