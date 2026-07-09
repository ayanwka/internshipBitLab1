package dtos

type CourseDTO struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ChapterDTO struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CourseID     uint   `json:"course_id"`
	ChapterOrder uint   `json:"chapter_order"`
}

type LessonDTO struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Content     string `json:"content"`
	ChapterID   uint   `json:"chapter_id"`
	LessonOrder uint   `json:"lesson_order"`
}
