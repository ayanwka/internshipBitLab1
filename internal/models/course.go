package models

import "time"

func (Course) TableName() string {
	return "course" // точное имя таблицы из миграции
}

type Course struct {
	Id          uint
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
