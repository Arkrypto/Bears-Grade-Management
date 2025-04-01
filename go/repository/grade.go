package repository

import (
	"StudentManagement/model"
	"fmt"
)

var Grades = []model.Grade{
	{1, 1, 60},
	{2, 1, 60},
	{3, 1, 60},
	{4, 1, 60},
	{5, 1, 60},
	{6, 1, 60},
	{7, 1, 60},
	{8, 1, 60},
	{9, 1, 60},
}

func GetGrades(id int) map[int]float64 {
	var grades = make(map[int]float64, 9)
	for _, grade := range Grades {
		if grade.StudentID == id {
			grades[grade.SubjectID] = grade.Score
		}
	}
	return grades
}

func Main() {
	fmt.Println("nmsl")
}
