package routers

import (
	h "inventory/handler"

	"github.com/gin-gonic/gin"
)

// ItemRoutes : router to set handler for item related api
func ItemRoutes(r *gin.RouterGroup) {
	r.GET("", h.GetItems)
	r.POST("", h.AddItem)
	r.PUT("/:id", h.ModifyItem)
	r.DELETE("/:id", h.RemoveItem)
}
