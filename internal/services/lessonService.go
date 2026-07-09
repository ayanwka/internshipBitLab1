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

type LessonService interface {
	CreateLesson(lesson dtos.LessonDTO) error
	GetLesson(id uint) (dtos.LessonDTO, error)
	GetLessonsByChapterID(chapterID uint) ([]dtos.LessonDTO, error)
	UpdateLesson(updLesson dtos.LessonDTO) error
	DeleteLesson(id uint) error
}

type lessonService struct {
	lessonRepository repositories.LessonRepository
}

func NewLessonService(lessonRepository repositories.LessonRepository) LessonService {
	return &lessonService{lessonRepository: lessonRepository}
}

func (l *lessonService) CreateLesson(lesson dtos.LessonDTO) error {
	if lesson.Name == "" || lesson.Content == "" {
		return apperrors.NewValidationError("имя или содержимое лекции не могут быть пустыми")
	}
	lessonModel := &models.Lesson{
		Name:        lesson.Name,
		Description: lesson.Description,
		Content:     lesson.Content,
		ChapterId:   lesson.ChapterID,
		LessonOrder: lesson.LessonOrder,
	}
	err := l.lessonRepository.CreateLesson(lessonModel)
	if err != nil {
		logrus.Errorf("Ошибка при создании новой лекции")
		return err
	}
	logrus.WithFields(logrus.Fields{
		"lesson_id":  lesson.Id,
		"name":       lesson.Name,
		"chapter_id": lesson.ChapterID,
	}).Debug("детали лекции")
	return nil
}

func (l *lessonService) GetLesson(id uint) (dtos.LessonDTO, error) {
	lesson, err := l.lessonRepository.GetLesson(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dtos.LessonDTO{}, apperrors.NewNotFoundError("Lesson", id)
		}
		logrus.Errorf("Ошибка при получении лекции")
		return dtos.LessonDTO{}, err
	}
	return mappers.MapToLessonDto(*lesson), nil
}

func (l *lessonService) GetLessonsByChapterID(chapterID uint) ([]dtos.LessonDTO, error) {
	_, err := l.lessonRepository.GetBy("id = ?", chapterID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NewNotFoundError("Chapter", chapterID)
		}
		return nil, err
	}
	lessons, err := l.lessonRepository.GetLessonsByChapterID(chapterID)
	if err != nil {
		logrus.Errorf("Ошибка при получении списка всех лекций текущей главы")
		return nil, err
	}
	return mappers.MapToDtoLessonList(*lessons), nil
}

func (l *lessonService) UpdateLesson(updLesson dtos.LessonDTO) error {
	if updLesson.Name == "" || updLesson.Content == "" {
		return apperrors.NewValidationError("имя или содержимое лекции не могут быть пустыми")
	}
	_, err := l.lessonRepository.GetBy("id = ?", updLesson.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.NewNotFoundError("Lesson", updLesson.Id)
		}
		logrus.Errorf("Ошибка при обновлении текущей лекции")
		return err
	}

	lessonModel := &models.Lesson{
		Id:          updLesson.Id,
		Name:        updLesson.Name,
		Description: updLesson.Description,
		Content:     updLesson.Content,
		ChapterId:   updLesson.ChapterID,
		LessonOrder: updLesson.LessonOrder,
	}
	err = l.lessonRepository.UpdateLesson(lessonModel)
	if err != nil {
		logrus.Errorf("ошибка при сохранении обновленных данных лекции: %v", err)
		return err
	}
	logrus.WithFields(logrus.Fields{
		"lesson_id":  lessonModel.Id,
		"name":       lessonModel.Name,
		"chapter_id": lessonModel.ChapterId,
	}).Debug("Детали обновленной лекции в БД")
	return nil
}

func (l *lessonService) DeleteLesson(id uint) error {
	_, err := l.lessonRepository.GetBy("id = ?", id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.NewNotFoundError("Lesson", id)
		}
		return err
	}
	err = l.lessonRepository.DeleteLesson(id)
	if err != nil {
		logrus.Errorf("ошибка при удалении лекции из БД: %v", err)
		return err
	}
	return nil
}
