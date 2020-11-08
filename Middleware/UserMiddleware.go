package Middleware

import (
	"CRUDTEST/Models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CreateUser Validator
func NameMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Models.User
		// c.BindJSON(&user)
		// c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(reqBody)))
		// fmt.Println(user.Rollnum)
		// if user.Name == "" {
		// 	c.JSON(http.StatusNotFound, gin.H{"No Key": "Name"})
		// 	c.AbortWithStatus(http.StatusNotFound)
		// }
		buf := make([]byte, 1024)
		num, _ := c.Request.Body.Read(buf)
		reqBody := string(buf[0:num])
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(reqBody))) // Write body back
		// fmt.Println(reqBody)
		json.Unmarshal([]byte(reqBody), &user)
		fmt.Println("Email ", user.Email)
		if user.Name == "" {
			c.JSON(http.StatusNotFound, gin.H{"No Key": "Name"})
			c.AbortWithStatus(http.StatusNotFound)
		}
	}

}

//Dummy middleware
// func EmailMiddleware(c *gin.Context) {
// 	var user Models.User
// 	c.BindJSON(&user)
// 	if user.Email == "" {
// 		c.JSON(http.StatusNotFound, gin.H{"No Key": "Email"})
// 		c.AbortWithStatus(http.StatusNotFound)
// 		return
// 	}
// 	c.Next()
// }

// //Dummy middleware
// func PhoneMiddleware(c *gin.Context) {
// 	var user Models.User
// 	c.BindJSON(&user)
// 	fmt.Println(user.Phone)
// 	if user.Phone == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"No Key": "Phone"})
// 		c.AbortWithStatus(http.StatusNotFound)
// 		return
// 	}
// 	// Pass on to the next-in-chain
// 	c.Next()
// }
