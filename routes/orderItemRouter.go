package routes

import (
	"github.com/gin-gonic/gin"
	controllers "restaurant-management/controllers"
)

func OrderItemRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incommingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItems)
	incommingRoutes.GET("orderItems-order/:orderItem_id", controllers.GetOrderItemsByOder())
	incommingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incommingRoutes.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())

}
