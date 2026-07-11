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

// @Summary      Creating a chapter
// @Param		 input body dtos.ChapterDTO true "Данные главы"
// @Failure		 400 {object} apperrors.ValidationError
// @Failure		 500 {object} map[string]string
// @Router		 /chapters [post]
// @Tags 		 Chapters
// @Success 	 201 {object} map[string]string
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
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "глава успешно создана"})
}

// @Summary      Show the chapter
// @Param		 id path uint true "ID главы"
// @Failure		 400 {object} apperrors.ValidationError
// @Failure		 500 {object} map[string]string
// @Router		 /chapters/{id} [get]
// @Tags 		 Chapters
// @Success 	 200 {object} map[string]string
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
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, chapter)
}

// @Summary      Show a list of chapters by course ID
// @Param		 id path uint true "ID курса"
// @Failure		 400 {object} apperrors.ValidationError
// @Failure		 500 {object} map[string]string
// @Router		 /courses/{id}/chapters [get]
// @Tags 		 Chapters
// @Success 	 200 {object} map[string]string
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
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, chapters)
}

// @Summary      Chapter update
// @Param		 id path uint true "ID главы"
// @Failure		 400 {object} apperrors.ValidationError
// @Failure		 500 {object} map[string]string
// @Router		 /chapters/{id} [put]
// @Tags 		 Chapters
// @Success 	 200 {object} map[string]string
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
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, chapter)
}

// @Summary      Delete a chapter
// @Param		 id path uint true "ID главы"
// @Failure		 400 {object} apperrors.ValidationError
// @Failure		 500 {object} map[string]string
// @Router		 /chapters/{id} [delete]
// @Tags 		 Chapters
// @Success 	 200 {object} map[string]string
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
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "глава успешно удалена"})
}
