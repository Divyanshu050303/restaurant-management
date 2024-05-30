package routes

import (
	controllers "restaurant-management/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/users", controllers.GetUsers())
	incommingRoutes.GET("/users/:user_id", controllers.GetUser())
	incommingRoutes.POST("/users/signup", controllers.SignUp())
	incommingRoutes.PATCH("/users/login", controllers.SignIn())

}
