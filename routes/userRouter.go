package routes

import (
	"github.com/gin-gonic/gin"
	controllers "restaurant-management/controllers"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/users", controllers.GetUsers())
	incommingRoutes.GET("/users/:user_id", controllers.GetUsers())
	incommingRoutes.POST("/users/signup", controllers.SignUp())
	incommingRoutes.PATCH("/users/login", controllers.SignIn())

}
