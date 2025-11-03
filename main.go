package main

import (
	"example/restapi/db"
	"example/restapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	//initiate Database
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run()
}
