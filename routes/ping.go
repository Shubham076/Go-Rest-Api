package routes

import (
	"BootCampT1/controllers"
	"github.com/gin-gonic/gin"
)

func PingRoutes(router *gin.Engine) {
	router.GET("/ping", controllers.DefaultHandler)
}
