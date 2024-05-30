package routes

import (
	controllers "restaurant-management/controller"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/foods", controllers.GetFoods())
	incommingRoutes.GET("/foods/:food_id", controllers.GetFood())
	incommingRoutes.GET("/foods", controllers.CreateFood())
	incommingRoutes.GET("/foods/:foods", controllers.UpdateFood())

}
