package router

import (
	"mini_project_restapi/controller"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()
	router.GET("/persons", controller.GetAllPerson)
	router.POST("/persons", controller.InsertPerson)
	router.PUT("/persons/:id", controller.UpdatePerson)
	router.DELETE("/persons/:id", controller.DeletePerson)

	router.Run(":8100")
}
