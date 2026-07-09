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

type ChapterService interface {
	CreateChapter(chapter dtos.ChapterDTO) error
	GetChapter(id uint) (dtos.ChapterDTO, error)
	GetChaptersByCourseID(courseID uint) ([]dtos.ChapterDTO, error)
	UpdateChapter(updChapter dtos.ChapterDTO) error
	DeleteChapter(id uint) error
}

type chapterService struct {
	chapterRepository repositories.ChapterRepository
}

func NewChapterService(chapterRepository repositories.ChapterRepository) ChapterService {
	return &chapterService{chapterRepository: chapterRepository}
}

func (c *chapterService) CreateChapter(chapter dtos.ChapterDTO) error {
	if chapter.Name == "" || chapter.Description == "" {
		return apperrors.NewValidationError("имя или описание главы не могут быть пустыми")
	}
	chapterModel := &models.Chapter{
		Name:         chapter.Name,
		Description:  chapter.Description,
		CourseId:     chapter.CourseID,
		ChapterOrder: chapter.ChapterOrder,
	}
	err := c.chapterRepository.CreateChapter(chapterModel)
	if err != nil {
		logrus.Errorf("Ошибка при сохранении главы в БД: %v", err)
		return err
	}
	logrus.WithFields(logrus.Fields{
		"chapter_id": chapter.Id,
		"name":       chapter.Name,
		"course_id":  chapter.CourseID,
	}).Debug("детали главы")
	return nil
}

func (c *chapterService) GetChapter(id uint) (dtos.ChapterDTO, error) {
	chapter, err := c.chapterRepository.GetChapter(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dtos.ChapterDTO{}, apperrors.NewNotFoundError("Chapter", id)
		}
		logrus.Errorf("Ошибка при получении главы: %v", err)
		return dtos.ChapterDTO{}, err
	}
	return mappers.MapToChapterDto(*chapter), nil
}

func (c *chapterService) GetChaptersByCourseID(courseID uint) ([]dtos.ChapterDTO, error) {
	_, err := c.chapterRepository.GetBy("id = ?", courseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NewNotFoundError("Course", courseID)
		}
		return nil, err
	}
	chapters, err := c.chapterRepository.GetChaptersByCourseID(courseID)
	if err != nil {
		logrus.Errorf("Ошибка при получении списка глав из текущего курса: %v", err)
		return nil, err
	}
	return mappers.MapToChapterDtoList(*chapters), nil
}

func (c *chapterService) UpdateChapter(updChapter dtos.ChapterDTO) error {
	if updChapter.Name == "" || updChapter.Description == "" {
		return apperrors.NewValidationError("имя или описание главы не могут быть пустыми")
	}
	_, err := c.chapterRepository.GetBy("id = ?", updChapter.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.NewNotFoundError("Chapter", updChapter.Id)
		}
		logrus.Errorf("ошибка при обновлении данных главы: %v", err)
		return err
	}

	chapterModel := &models.Chapter{
		Id:           updChapter.Id,
		Name:         updChapter.Name,
		Description:  updChapter.Description,
		CourseId:     updChapter.CourseID,
		ChapterOrder: updChapter.ChapterOrder,
	}
	err = c.chapterRepository.UpdateChapter(chapterModel)
	if err != nil {
		logrus.Errorf("ошибка при сохранении обновленных данных главы: %v", err)
		return err
	}
	logrus.WithFields(logrus.Fields{
		"chapter_id": updChapter.Id,
		"name":       updChapter.Name,
		"course_id":  updChapter.CourseID,
	}).Debug("Детали обновленной главы в БД")
	return nil
}

func (c *chapterService) DeleteChapter(id uint) error {
	_, err := c.chapterRepository.GetBy("id = ?", id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.NewNotFoundError("Chapter", id)
		}
		return err
	}
	err = c.chapterRepository.DeleteChapter(id)
	if err != nil {
		logrus.Errorf("ошибка при удалении главы из БД: %v", err)
		return err
	}
	return nil
}
