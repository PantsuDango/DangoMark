package routes

import (
	"DangoMark/service/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	router := gin.Default()
	router.Use(cors.Default())

	Controller := new(controller.Controller)
	router.POST("/DangoMark/api", Controller.Handle)

	return router
}
