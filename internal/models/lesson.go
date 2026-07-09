package models

import "time"

func (Lesson) TableName() string {
	return "lesson" // точное имя таблицы из миграции
}

type Lesson struct {
	Id          uint
	Name        string
	Description string
	Content     string
	LessonOrder uint
	ChapterId   uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
