package controllers

import (
	"fmt"
	"net/http"

	"github.com/bilalshabbir31/bun_learning/common"
	"github.com/bilalshabbir31/bun_learning/models"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)


var DB *bun.DB

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

func Show_all(ctx *gin.Context) ([]common.Teacher, error) {
	teachers, err := models.Fetch_all_teacher(ctx,DB)
	if err == nil {
		return teachers, nil
	} else {
		return nil, fmt.Errorf("error controller get all to dos task %v", err)
	}
}