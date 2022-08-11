package main

import (
	"github.com/anto1290/go-dts/controllers/listcontroller"

	"github.com/anto1290/go-dts/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())
	models.ConnectionDatabase()
	// Api
	r.GET("api/list", listcontroller.GetAll)
	r.GET("api/list/:id", listcontroller.Get)
	r.POST("api/list", listcontroller.Create)
	r.PUT("api/list/:id", listcontroller.Update)
	r.DELETE("api/list/:id", listcontroller.Delete)
	// views

	r.Run()
}
