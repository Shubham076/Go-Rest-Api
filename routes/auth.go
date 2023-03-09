package routes

import (
	"BootCampT1/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/login", controllers.LoginHandler)
	router.POST("/logout", controllers.Logouthandler)
	router.POST("/signUp", controllers.SignUpHandler)
	router.POST("/verifyOtp", controllers.VerifyOtp)
	router.GET("/user", controllers.GetUserData)
}
