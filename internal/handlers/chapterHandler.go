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

type ChapterHandler struct {
	chapterService services.ChapterService
}

func NewChapterHandler(chapterService services.ChapterService) *ChapterHandler {
	return &ChapterHandler{chapterService: chapterService}
}

func (c2 *ChapterHandler) CreateChapter(c *gin.Context) {
	var chapter dtos.ChapterDTO
	if err := c.ShouldBindJSON(&chapter); err != nil {
		_ = c.Error(apperrors.NewValidationError("некорректный формат главы"))
		c.Abort()
		return
	}
	logrus.Info("Создание новой главы")
	err := c2.chapterService.CreateChapter(chapter)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "глава успешно создана"})
}

func (c2 *ChapterHandler) GetChapter(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		_ = c.Error(apperrors.NewValidationError("некорректный ID главы"))
		c.Abort()
		return
	}
	logrus.Info("Получение текущей главы")
	chapter, err := c2.chapterService.GetChapter(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, chapter)
}

func (c2 *ChapterHandler) GetChaptersByCourseID(c *gin.Context) {
	idStr := c.Param("id")
	CourseID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		_ = c.Error(apperrors.NewValidationError("некорректный ID курса"))
		c.Abort()
		return
	}
	logrus.Info("Получение полного списка глав текущего курса")
	chapters, err := c2.chapterService.GetChaptersByCourseID(uint(CourseID))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, chapters)
}

func (c2 *ChapterHandler) UpdateChapter(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		_ = c.Error(apperrors.NewValidationError("некорректный ID главы"))
		c.Abort()
		return
	}
	var chapter dtos.ChapterDTO
	if err = c.ShouldBindJSON(&chapter); err != nil {
		_ = c.Error(apperrors.NewValidationError("некорректный формат главы"))
		c.Abort()
		return
	}
	logrus.Info("Обновление текущей главы")
	chapter.Id = uint(id)
	err = c2.chapterService.UpdateChapter(chapter)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, chapter)
}

func (c2 *ChapterHandler) DeleteChapter(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		_ = c.Error(apperrors.NewValidationError("некорректный ID главы"))
		c.Abort()
		return
	}
	logrus.Info("Удаление текущей главы")
	err = c2.chapterService.DeleteChapter(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "глава успешно удалена"})
}
