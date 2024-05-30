package routes

import (
	controllers "restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.Get("/foods", controllers.GetFoods())
	incommingRoutes.Get("/foods/:food_id", controllers.GetFood())
	incommingRoutes.Get("/foods", controllers.CreateFood())
	incommingRoutes.Get("/foods/:foods", controllers.UpdateFood())

}
