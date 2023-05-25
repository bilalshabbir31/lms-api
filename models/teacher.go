package models

import (
	"context"
	"errors"
	"fmt"

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

func GetTeacherByID(ctx context.Context, db *bun.DB, id int) (common.Teacher, error) {
	teacher := &Teacher{ID: id}
	err := db.NewSelect().Model(teacher).WherePK().Scan(ctx)
	if err != nil {
		return common.Teacher{}, fmt.Errorf("failed to get teacher: %w", err)
	}

	teacherRes := common.Teacher{
		ID:   teacher.ID,
		Name: teacher.Name,
	}

	fmt.Println(teacher.ID, teacher.Name)

	return teacherRes, nil
}

func Delete_Teacher(ctx context.Context, db *bun.DB,id int)(bool,error){
	teacher := &Teacher{ID: id}
	res, err := db.NewDelete().Model(teacher).WherePK().Exec(ctx)
	if err != nil {
		return false, fmt.Errorf("error model delete teacher %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("error getting rows affetced: %w", err)
	}

	if rowsAffected == 0 {
		return false, errors.New("nothing update or teachet not found")
	}

	return true, nil
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