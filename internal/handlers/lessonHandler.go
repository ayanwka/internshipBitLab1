package handlers

import (
	"lms-main-service/internal/dtos"
	"lms-main-service/internal/services"
	"lms-main-service/pkg/apperrors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LessonHandler struct {
	lessonService services.LessonService
}

func NewLessonHandler(lessonService services.LessonService) *LessonHandler {
	return &LessonHandler{lessonService: lessonService}
}

func (l *LessonHandler) CreateLesson(c *gin.Context) {
	var lesson dtos.LessonDTO
	if err := c.ShouldBindJSON(&lesson); err != nil {
		_ = c.Error(apperrors.NewValidationError("Некорректный формат лекции"))
		c.Abort()
		return
	}
	logrus.Info("Создание новой лекции")
	err := l.lessonService.CreateLesson(lesson)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "лекция успешно создана"})
}

func (l *LessonHandler) GetLesson(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		_ = c.Error(apperrors.NewValidationError("некорректный ID лекции"))
		c.Abort()
		return
	}
	logrus.Info("Получение текущей лекции")
	lesson, err := l.lessonService.GetLesson(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, lesson)
}

func (l *LessonHandler) GetLessonsByChapterID(c *gin.Context) {
	idStr := c.Param("id")
	chapterID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		_ = c.Error(apperrors.NewValidationError("некорректный ID главы"))
		c.Abort()
		return
	}
	logrus.Info("Получение списка всех лекций текущей главы")
	lessons, err := l.lessonService.GetLessonsByChapterID(uint(chapterID))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, lessons)
}

func (l *LessonHandler) UpdateLesson(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		_ = c.Error(apperrors.NewValidationError("некорректный ID лекции"))
		c.Abort()
		return
	}
	var lessons dtos.LessonDTO
	err = c.ShouldBindJSON(&lessons)
	if err != nil {
		_ = c.Error(apperrors.NewValidationError("некорректный формат лекции"))
		c.Abort()
		return
	}
	logrus.Info("Обновление текущей лекции")
	lessons.Id = uint(id)
	err = l.lessonService.UpdateLesson(lessons)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, lessons)
}

func (l *LessonHandler) DeleteLesson(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		_ = c.Error(apperrors.NewValidationError("некорректный ID лекции"))
		c.Abort()
		return
	}
	logrus.Info("Удаление текущей лекции")
	err = l.lessonService.DeleteLesson(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "лекция успешно удалена"})
}
