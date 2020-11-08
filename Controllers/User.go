package Controllers

import (
	"CRUDTEST/Models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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
	c.BindJSON(&user)
	fmt.Println("ROLLL: " + user.Rollnum)
	fmt.Println(user)
	err := Models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), "Duplicate") {
			c.JSON(http.StatusNotFound, gin.H{"Duplicate Key": " Roll Number"})
		} else if strings.Contains(err.Error(), "1364") {
			if strings.Contains(err.Error(), "rollnum") {
				c.JSON(http.StatusNotFound, gin.H{"Missing Key": " Roll Number"})
			}
		}
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
	return
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"id " + id: "User not found"})
	} else {
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
