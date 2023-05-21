package main

import (
	"github.com/bilalshabbir31/bun_learning/connect"
	"github.com/bilalshabbir31/bun_learning/models"
)



func main(){
	db := connect.Connect()
	

	// models.Create_teacher(db,4,"Obaid ur Rehman")
	models.Fetch_all_teacher(db)
	models.Get_teacher_by_id(db,1)
}