package main

import (
	"BootCampT1/config"
	"BootCampT1/external/RabbitMq"
	"BootCampT1/external/rds"
	"BootCampT1/external/ses"
	"BootCampT1/logger"
	"BootCampT1/routes"
	"github.com/gin-gonic/gin"
)

func enableRoutes(r *gin.Engine) {
	routes.AuthRoutes(r)
	routes.PingRoutes(r)
}

func startServer(port string) {
	r := gin.Default()
	enableRoutes(r)
	r.Run(port)
}
func main() {
	logger.Init()
	conf, err := config.LoadConfig()
	if err != nil {
		logger.Error.Println("Error during loading config file: %v", err)
	}
	rds.Connect()
	rds.CreateTables()
	RabbitMq.Init()
	ses.Init()
	startServer(conf.App.Port)
}
