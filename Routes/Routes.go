package Routes

import (
	"CRUDTEST/Controllers"
	"CRUDTEST/Middleware"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Router Setup
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	grp1.POST("login", Controllers.Login)
	grp1.GET("users", Controllers.GetUsers)
	grp1.POST("user", Middleware.NameMiddleware(), Controllers.CreateUser)
	grp1.GET("user", Middleware.CheckMiddleware(), Controllers.GetUserByID)
	grp1.PUT("user/:id", Controllers.UpdateUser)
	grp1.DELETE("user/:id", Controllers.DeleteUser)
	return r
}
