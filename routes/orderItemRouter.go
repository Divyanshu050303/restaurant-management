package routes

import (
	"github.com/gin-gonic/gin"
	controllers "restaurant-management/controller"
)

func OrderItemRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incommingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	incommingRoutes.GET("orderItems-order/:orderItem_id", controllers.GetOrderItemsByOrder())
	incommingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incommingRoutes.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())

}
