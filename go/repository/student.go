package repository

import (
	"StudentManagement/model"
)

var students = []model.Student{
	{1, "熊熊", 1},
	{2, "球吊", 1},
	{3, "TM", 1},
	{4, "戴某", 1},
	{5, "兴根", 1},
	{6, "郭某", 1},
	{7, "强哥", 1},
	{8, "彭奇", 1},
}

func GetStudents() []model.Student {
	return students
}

func GetStudentsWithGrades() []model.StudentWithGrades {
	var studentsWithGrades = make([]model.StudentWithGrades, len(students))
	for i := 0; i < len(students); i++ {
		studentsWithGrades[i] = model.StudentWithGrades{
			ID:     students[i].ID,
			Name:   students[i].Name,
			Gender: students[i].Gender,
			Grades: GetGrades(students[i].ID),
		}
	}
	return studentsWithGrades
}

func AddStudent(student model.Student) bool {
	student.ID = len(students) + 1
	students = append(students, student)
	return true
}

func GetStudentById(id int) model.Student {
	for _, student := range students {
		if student.ID == id {
			return student
		}
	}
	return model.NilStudent()
}

func GetStudentByName(name string) model.Student {
	for _, student := range students {
		if student.Name == name {
			return student
		}
	}
	return model.NilStudent()
}

func DeleteStudentById(id int) bool {
	for index, student := range students {
		if student.ID == id {
			students = append(students[:index], students[index+1:]...)
			return true
		}
	}
	return false
}
