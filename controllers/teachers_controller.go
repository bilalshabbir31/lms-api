package controllers

import (
	"fmt"
	"net/http"

	"github.com/bilalshabbir31/bun_learning/common"
	"github.com/bilalshabbir31/bun_learning/models"
	"github.com/gin-gonic/gin"
)



func Create(ctx *gin.Context) (int,error){
	var teacher common.Teacher
	err := ctx.ShouldBind(&teacher)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error": err.Error(),
		})
		return -1,err
	}
	teacher_id,err:=models.Create_teacher(ctx,DB,teacher.ID,teacher.Name)
	if err==nil{
		return teacher_id,nil
	}else{
		return -1, fmt.Errorf("error in controller to create a teacher %v",err )
	}
}