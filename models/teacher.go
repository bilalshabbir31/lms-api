package models

import (
	"context"

	"github.com/uptrace/bun"
)

type Teacher struct {
	bun.BaseModel `bun:"teachers"`
	ID            int    `bun:"id,pk"`
	Name          string `bun:"name"`
}

func Create_teacher(db *bun.DB, id int, name string) (int, error) {
	ctx := context.Background()
	teacher := &Teacher{ID: id, Name: name}
	_, err := db.NewInsert().Model(teacher).Exec(ctx)
	if err != nil {
		println(err)
	}

	return teacher.ID, err
}

func Fetch_all_teacher(db *bun.DB){
	ctx := context.Background()
	var teachers []Teacher
	err:= db.NewSelect().Model(&teachers).Scan(ctx)

	if err!=nil{
		println(err)
	}
	for _, teacher := range teachers {
		print(teacher.ID)
		print(teacher.Name)
		println()
	}
}

func Get_teacher_by_id(db *bun.DB,id int) (int,error){

	ctx := context.Background()

	var teacher Teacher
	err:= db.NewSelect().Model(&teacher).Where("id=?",id).Scan(ctx)

	if err!=nil{
		println(err)
	}

	println(teacher.ID,teacher.Name)
	return teacher.ID,err

}
