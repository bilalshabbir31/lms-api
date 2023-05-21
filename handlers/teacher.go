package handlers

import (
	"net/http"

	"github.com/bilalshabbir31/bun_learning/controllers"
	"github.com/gin-gonic/gin"
)




func Create_teacher(ctx *gin.Context)  {
	teacher, err := controllers.Create(ctx)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{ "error in handler to create teacher": err.Error() })
	}else{
		ctx.JSON(http.StatusOK,teacher)
	}
}

func Show_all_teachers(ctx *gin.Context)  {
	teachers, err := controllers.Show_all(ctx)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{" error in show handler ": err.Error})
	}else{
		ctx.JSON(http.StatusOK,teachers)
	}
	
}