package routes

import (
	"github.com/gin-gonic/gin"
	controllers "restaurant-management/controllers"
)

func OrderRoutes(incommingRoutes *gin.Engine) {

	incommingRoutes.GET("/orders", controllers.Getorders())
	incommingRoutes.GET("/orders/:order_id", controllers.Getorder())
	incommingRoutes.POST("/orders", controllers.Createorder())
	incommingRoutes.PATCH("/orders/:order_id", controllers.Updateorder())

}
