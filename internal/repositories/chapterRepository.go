package repositories

import (
	"lms-main-service/internal/models"

	"gorm.io/gorm"
)

type ChapterRepository interface {
	CreateChapter(chapter *models.Chapter) error
	GetChapter(id uint) (*models.Chapter, error)
	GetChaptersByCourseID(courseID uint) (*[]models.Chapter, error)
	UpdateChapter(updChapter *models.Chapter) error
	DeleteChapter(id uint) error
	GetBy(query string, args ...any) (*models.Chapter, error)
}

type chapterRepository struct {
	db *gorm.DB
}

func NewChapterRepository(db *gorm.DB) ChapterRepository {
	return &chapterRepository{db: db}
}

func (c *chapterRepository) CreateChapter(chapter *models.Chapter) error {
	err := c.db.Create(chapter).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *chapterRepository) GetChapter(id uint) (*models.Chapter, error) {
	var chapter models.Chapter
	err := c.db.First(&chapter, id).Error
	if err != nil {
		return nil, err
	}
	return &chapter, nil
}

func (c *chapterRepository) GetChaptersByCourseID(courseID uint) (*[]models.Chapter, error) {
	var chapters []models.Chapter
	err := c.db.Where("course_id = ?", courseID).Order("chapter_order").Find(&chapters).Error
	if err != nil {
		return nil, err
	}
	return &chapters, nil
}

func (c *chapterRepository) UpdateChapter(updChapter *models.Chapter) error {
	err := c.db.Save(updChapter).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *chapterRepository) DeleteChapter(id uint) error {
	err := c.db.Delete(&models.Chapter{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *chapterRepository) GetBy(query string, args ...any) (*models.Chapter, error) {
	var chapter models.Chapter
	// Ищем по полю query. В GORM это делается через Where или явное указание структуры
	// GORM подставит query (например, "id = ?") и args (например, 5)
	err := c.db.Where(query, args...).First(&chapter).Error
	if err != nil {
		return nil, err // Если не нашли или упала база — возвращаем ошибку
	}
	return &chapter, nil // Возвращаем указатель на главу
}
