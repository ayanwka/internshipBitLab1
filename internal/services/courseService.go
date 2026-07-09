package services

import (
	"errors"
	"lms-main-service/internal/dtos"
	"lms-main-service/internal/mappers"
	"lms-main-service/internal/models"
	"lms-main-service/internal/repositories"
	"lms-main-service/pkg/apperrors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CourseService interface {
	CreateCourse(course dtos.CourseDTO) error
	GetCourse(id uint) (dtos.CourseDTO, error)
	GetAllCourses() ([]dtos.CourseDTO, error)
	UpdateCourse(updCourse dtos.CourseDTO) error
	DeleteCourse(id uint) error
}

type courseService struct {
	courseRepository repositories.CourseRepository
}

func NewCourseService(courseRepository repositories.CourseRepository) CourseService {
	return &courseService{courseRepository: courseRepository}
}

func (c *courseService) CreateCourse(course dtos.CourseDTO) error {
	if course.Name == "" || course.Description == "" {
		return apperrors.NewValidationError("имя или описание курса не могут быть пустыми")
	}
	courseModel := &models.Course{
		Name:        course.Name,
		Description: course.Description,
	}
	_, err := c.courseRepository.GetBy("name = ?", course.Name)
	if err == nil {
		return apperrors.NewConflictError("имя с таким курсом уже существует")
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		err = c.courseRepository.CreateCourse(courseModel)
		if err != nil {
			logrus.Errorf("ошибка при сохранении курса в БД: %v", err)
			return err
		}
		logrus.WithFields(logrus.Fields{
			"course_id": course.Id,
			"name":      course.Name,
		}).Debug("детали курса")
		return nil
	} else {
		logrus.Errorf("ошибка при создании курса: %v", err)
		return err
	}
}

func (c *courseService) GetCourse(id uint) (dtos.CourseDTO, error) {
	course, err := c.courseRepository.GetCourse(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dtos.CourseDTO{}, apperrors.NewNotFoundError("Course", id)
		}
		logrus.Errorf("ошибка при получении курса: %v", err)
		return dtos.CourseDTO{}, err
	}
	return mappers.MapToCourseDto(*course), nil
}

func (c *courseService) GetAllCourses() ([]dtos.CourseDTO, error) {
	courses, err := c.courseRepository.GetAllCourses()
	if err != nil {
		logrus.Errorf("ошибка при выводе всех курсов: %v", err)
		return nil, err
	}
	return mappers.MapToCourseDtoList(*courses), nil
}

func (c *courseService) UpdateCourse(updCourse dtos.CourseDTO) error {
	if updCourse.Name == "" || updCourse.Description == "" {
		return apperrors.NewValidationError("имя или описание курса не могут быть пустыми")
	}
	_, err := c.courseRepository.GetBy("id = ?", updCourse.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.NewNotFoundError("Course", updCourse.Id)
		}
		logrus.Errorf("ошибка при обновлении данных курса: %v", err)
		return err
	}

	courseModel := &models.Course{
		Id:          updCourse.Id,
		Name:        updCourse.Name,
		Description: updCourse.Description,
	}
	err = c.courseRepository.UpdateCourse(courseModel)
	if err != nil {
		logrus.Errorf("ошибка при сохранении обновленных данных курса: %v", err)
		return err
	}
	logrus.WithFields(logrus.Fields{
		"course_id": updCourse.Id,
		"name":      updCourse.Name,
	}).Debug("Детали обновленного курса в БД")
	return nil
}

func (c *courseService) DeleteCourse(id uint) error {
	_, err := c.courseRepository.GetBy("id = ?", id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.NewNotFoundError("Course", id)
		}
	}
	err = c.courseRepository.DeleteCourse(id)
	if err != nil {
		logrus.Errorf("ошибка при удалении курса из БД: %v", err)
		return err
	}
	return nil
}
