package repositories

import (
	"lms-main-service/internal/models"

	"gorm.io/gorm"
)

type LessonRepository interface {
	CreateLesson(lesson *models.Lesson) error
	GetLesson(id uint) (*models.Lesson, error)
	GetLessonsByChapterID(chapterID uint) (*[]models.Lesson, error)
	UpdateLesson(updLesson *models.Lesson) error
	DeleteLesson(id uint) error
	GetBy(query string, args ...any) (*models.Lesson, error)
}

type lessonRepository struct {
	db *gorm.DB
}

func NewLessonRepository(db *gorm.DB) LessonRepository {
	return &lessonRepository{db: db}
}

func (l *lessonRepository) CreateLesson(lesson *models.Lesson) error {
	err := l.db.Create(lesson).Error
	if err != nil {
		return err
	}
	return nil
}

func (l *lessonRepository) GetLesson(id uint) (*models.Lesson, error) {
	var lesson models.Lesson
	err := l.db.First(&lesson, id).Error
	if err != nil {
		return nil, err
	}
	return &lesson, nil
}

func (l *lessonRepository) GetLessonsByChapterID(chapterID uint) (*[]models.Lesson, error) {
	var lessons []models.Lesson
	err := l.db.Where("chapter_id = ?", chapterID).Order("lesson_order").Find(&lessons).Error
	if err != nil {
		return nil, err
	}
	return &lessons, nil
}

func (l *lessonRepository) UpdateLesson(updLesson *models.Lesson) error {
	err := l.db.Save(updLesson).Error
	if err != nil {
		return err
	}
	return nil
}

func (l *lessonRepository) DeleteLesson(id uint) error {
	err := l.db.Delete(&models.Lesson{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *lessonRepository) GetBy(query string, args ...any) (*models.Lesson, error) {
	var lesson models.Lesson
	// Ищем по полю query. В GORM это делается через Where или явное указание структуры
	// GORM подставит query (например, "id = ?") и args (например, 5)
	err := c.db.Where(query, args...).First(&lesson).Error
	if err != nil {
		return nil, err // Если не нашли или упала база — возвращаем ошибку
	}
	return &lesson, nil // Возвращаем указатель на курс
}
