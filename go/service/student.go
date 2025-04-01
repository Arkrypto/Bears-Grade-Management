package service

import (
	"StudentManagement/model"
	"StudentManagement/repository"
	"sort"
)

// 按 subject 序返回前 count 位学生名单
// subject: 0, 1, 2, 3 (分别表示总分、三科、理科、文科)

func GetSortedStudents(subject, count int) []model.StudentWithSubjectTag {
	var students = repository.GetStudentsWithGrades()
	switch subject {
	case 0:
		sort.Slice(students, func(i, j int) bool {
			return model.GetSum(students[i].Grades) > model.GetSum(students[j].Grades)
		})
	case 1:
		sort.Slice(students, func(i, j int) bool {
			return model.GetMainSum(students[i].Grades) > model.GetMainSum(students[j].Grades)
		})
	case 2:
		sort.Slice(students, func(i, j int) bool {
			return model.GetSciSum(students[i].Grades) > model.GetSciSum(students[j].Grades)
		})
	case 3:
		sort.Slice(students, func(i, j int) bool {
			return model.GetLibSum(students[i].Grades) > model.GetLibSum(students[j].Grades)
		})
	}

	var result = make([]model.StudentWithSubjectTag, len(students))
	for i := 0; i < len(students); i++ {
		result[i] = model.StudentWithSubjectTag{
			ID:     students[i].ID,
			Name:   students[i].Name,
			Gender: students[i].Gender,
			Grades: model.GradesWithSubjectTag(students[i].Grades),
		}
	}

	if count == 0 {
		return result
	}
	return result[0:count]
}

// 统计各科最高分

func GetTopStudents() map[string]model.StudentWithSubjectTag {
	var students = repository.GetStudentsWithGrades()
	var tops = make(map[int]model.StudentWithGrades, 9)
	for key, _ := range model.Subjects {
		tops[key] = model.NilStudentWithGrades()
	}
	// 遍历全部学生
	for _, student := range students {
		// 遍历所有科目，若当前学生当前科目大于记录的最高分，则更新该科目最高分学生为当前学生
		for subject, grade := range student.Grades {
			if grade > tops[subject].Grades[subject] {
				tops[subject] = student
			}
		}
	}
	// 将课程编号转化为课程名返回
	var result = make(map[string]model.StudentWithSubjectTag)
	for subject, student := range tops {
		result[model.Subjects[subject]] = model.StudentWithSubjectTag{
			ID:     student.ID,
			Name:   student.Name,
			Gender: student.Gender,
			Grades: model.GradesWithSubjectTag(student.Grades),
		}
	}
	return result
}
