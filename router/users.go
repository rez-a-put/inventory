package routers

import (
	h "inventory/handler"

	"github.com/gin-gonic/gin"
)

// UserRoutes : router to set handler for user related api
func UserRoutes(r *gin.RouterGroup) {
	r.POST("/login", h.Login)
}
