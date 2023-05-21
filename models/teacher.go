package models

import (
	"context"

	"github.com/bilalshabbir31/bun_learning/common"
	"github.com/uptrace/bun"
)

type Teacher struct {
	bun.BaseModel `bun:"teachers"`
	ID            int    `bun:"id,pk"`
	Name          string `bun:"name"`
}

func Create_teacher(ctx context.Context,db *bun.DB, id int, name string) (int, error) {
	teacher := &Teacher{ID: id, Name: name}
	_, err := db.NewInsert().Model(teacher).Exec(ctx)
	if err != nil {
		println(err)
	}

	return teacher.ID, err
}

func Fetch_all_teacher(ctx context.Context,db *bun.DB)  ([]common.Teacher, error) {
	var teachers []Teacher
	err:= db.NewSelect().Model(&teachers).Scan(ctx)

	if err!=nil{
		println(err)
	}
	return convert_resultset_of_teacher(teachers),nil
}

func Get_teacher_by_id(ctx context.Context,db *bun.DB,id int) (int,error){

	var teacher Teacher
	err:= db.NewSelect().Model(&teacher).Where("id=?",id).Scan(ctx)

	if err!=nil{
		println(err)
	}

	println(teacher.ID,teacher.Name)
	return teacher.ID,err

}


func convert_resultset_of_teacher(teachers []Teacher) []common.Teacher {
	result := make([]common.Teacher, len(teachers))

	for idx, teacher := range teachers {
		convert_teacher := common.Teacher{
			ID: teacher.ID,
			Name: teacher.Name,
		}
		result[idx] = convert_teacher
	}
	return result
}