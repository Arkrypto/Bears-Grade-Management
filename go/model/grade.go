package model

type Grade struct {
	SubjectID int
	StudentID int
	Score     float64
}

var Subjects = map[int]string{
	1: "语文",
	2: "数学",
	3: "英语",
	4: "物理",
	5: "化学",
	6: "生物",
	7: "政治",
	8: "历史",
	9: "地理",
}
