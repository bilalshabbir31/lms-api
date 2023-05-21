package main

import (
	"github.com/bilalshabbir31/bun_learning/connect"
	"github.com/bilalshabbir31/bun_learning/controllers"
	"github.com/bilalshabbir31/bun_learning/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)



func main(){
	db := connect.Connect()
	controllers.DB=db

	r := gin.Default()
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowAllOrigins= true
  r.Use(cors.New(ginConfig))

	r.POST("/create/teacher",handlers.Create_teacher)
	r.GET("/teachers",handlers.Show_all_teachers)


	r.Run()

}