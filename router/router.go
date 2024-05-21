package routers

import (
	"inventory/middleware"
	u "inventory/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRoutes : hold list of routes
func InitRoutes() {
	gf := gin.Default()

	rg := gf.Group("/inventory")

	// items
	itm := rg.Group("/items")
	itm.Use(middleware.JwtAuthMiddleware())
	ItemRoutes(itm)

	// users
	usr := rg.Group("/users")
	UserRoutes(usr)

	log.Fatal(http.ListenAndServe(u.GetEnvByKey("SERVER_HOST"), gf))
}
