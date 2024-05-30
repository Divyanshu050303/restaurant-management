package routes

import (
	"github.com/gin-gonic/gin"
	controllers "restaurant-management/controller"
)

func InvoiceRoutes(incommingRoutes *gin.Engine) {

	incommingRoutes.GET("/invoices", controllers.GetInvoices())
	incommingRoutes.GET("/invoices/:invoice_id", controllers.GetInvoice())
	incommingRoutes.POST("/invoices", controllers.CreateInvoice())
	incommingRoutes.PATCH("/invoices/:invoice_id", controllers.UpdateInvoice())

}
