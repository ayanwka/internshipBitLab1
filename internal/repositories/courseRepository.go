package repositories

import (
	"lms-main-service/internal/models"

	"gorm.io/gorm"
)

type CourseRepository interface {
	CreateCourse(course *models.Course) error
	GetCourse(id uint) (*models.Course, error)
	GetAllCourses() (*[]models.Course, error)
	UpdateCourse(updCourse *models.Course) error
	DeleteCourse(id uint) error
	GetBy(query string, args ...any) (*models.Course, error)
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (c *courseRepository) CreateCourse(course *models.Course) error {
	err := c.db.Create(course).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *courseRepository) GetCourse(id uint) (*models.Course, error) {
	var course models.Course
	err := c.db.First(&course, id).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (c *courseRepository) GetAllCourses() (*[]models.Course, error) {
	var courses []models.Course
	err := c.db.Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return &courses, nil
}

func (c *courseRepository) UpdateCourse(updCourse *models.Course) error {
	err := c.db.Save(updCourse).Error // обновляет всю строку
	if err != nil {
		return err
	}
	return nil
	// 2 вариант -- если надо обновить выборочно определенное поле, то используется --> c.db.Model(updCourse).Updates(updCourse)
}

func (c *courseRepository) DeleteCourse(id uint) error {
	err := c.db.Delete(&models.Course{}, id).Error
	// 2 вариант c.db.Model(&models.Course{}).Delete(id).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *courseRepository) GetBy(query string, args ...any) (*models.Course, error) {
	var course models.Course
	// Ищем по полю query. В GORM это делается через Where или явное указание структуры
	// GORM подставит query (например, "id = ?") и args (например, 5)
	err := c.db.Where(query, args...).First(&course).Error
	if err != nil {
		return nil, err // Если не нашли или упала база — возвращаем ошибку
	}
	return &course, nil // Нашли! Возвращаем указатель на курс
}
