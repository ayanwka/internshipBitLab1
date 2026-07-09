package models

import "time"

func (Chapter) TableName() string {
	return "chapter"
}

type Chapter struct {
	Id           uint
	Name         string
	Description  string
	ChapterOrder uint
	CourseId     uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
