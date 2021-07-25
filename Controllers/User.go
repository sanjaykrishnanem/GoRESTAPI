package Controllers

import (
	"CRUDTEST/Models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user Models.User
	// if err := c.ShouldBindWith(&user, binding.FormMultipart); err != nil {
	// 	log.Printf("%+v", err)
	// }
	c.Request.ParseForm()
	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	user.Phone, _ = strconv.ParseUint(c.PostForm("phone"), 10, 64)
	user.PasswordHash = c.PostForm("password")

	// fmt.Println("ROLLL: " + user.Rollnum)
	fmt.Println(user)
	err := Models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		// if strings.Contains(err.Error(), "Duplicate") {
		// 	c.JSON(http.StatusNotFound, gin.H{"Duplicate Key": " Roll Number"})
		// } else if strings.Contains(err.Error(), "1364") {
		// 	if strings.Contains(err.Error(), "email") {
		// 		c.JSON(http.StatusNotFound, gin.H{"Missing Key": " Email"})
		// 	}
		// }
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
	return
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	var uLogin struct {
		Email string `json:"email"`
	}
	c.ShouldBindBodyWith(&uLogin, binding.JSON)
	fmt.Println("Email =" + uLogin.Email)
	// id := c.Params.ByName("email")
	var user Models.User
	err := Models.GetUserByID(&user, uLogin.Email)
	if err != nil {
		// panic(err)
		c.JSON(http.StatusNotFound, gin.H{"email " + uLogin.Email: "User not found"})
	} else {
		fmt.Println("FOUND USER!!!!!!!!")
		c.JSON(http.StatusOK, user)
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"id " + id: "User not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"id " + id: "is deleted"})
	}
}

//Login ... Login Validation
func Login(c *gin.Context) {
	var uLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// c.ShouldBindBodyWith(&uLogin, binding.JSON)

	c.Request.ParseForm()
	uLogin.Email = c.PostForm("email")
	uLogin.Password = c.PostForm("password")

	err := Models.ValidateLogin(uLogin.Email, uLogin.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{uLogin.Email: "Login Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"email id: " + uLogin.Email: "has been logged in"})

}
