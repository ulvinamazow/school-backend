package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ulvinamazow/studentAPI/config"
	"github.com/ulvinamazow/studentAPI/routes"
)

func main() {
	r := gin.New()
	config.ConnnectDatabase()
	routes.UserRoute(r)
	r.Run(":4444")
}
