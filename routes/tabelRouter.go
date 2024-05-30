package routes

import (
	controllers "restaurant-management/controller"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incommingRoutes *gin.Engine) {

	incommingRoutes.GET("/tables", controllers.GetTables())
	incommingRoutes.GET("/tables/:table_id", controllers.GetTable())
	incommingRoutes.POST("/tables", controllers.CreateTable())
	incommingRoutes.PATCH("/tables/:table_id", controllers.UpdateTable())

}
