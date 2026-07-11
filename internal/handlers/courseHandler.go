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

type CourseHandler struct {
	courseService services.CourseService
}

func NewCourseHandler(courseService services.CourseService) *CourseHandler {
	return &CourseHandler{courseService: courseService}
}

// @Summary      Creating a course
// @Param		 input body dtos.CourseDTO true "Данные курса"
// @Failure		 400 {object} apperrors.ValidationError
// @Failure		 500 {object} map[string]string
// @Router		 /courses [post]
// @Tags 		 Courses
// @Success 	 201 {object} map[string]string
func (c2 *CourseHandler) CreateCourse(c *gin.Context) {
	var dto dtos.CourseDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		_ = c.Error(apperrors.NewValidationError("некорректный формат запроса"))
		c.Abort()
		return
	}
	logrus.Info("Создание нового курса")
	err := c2.courseService.CreateCourse(dto)
	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "курс успешно создан"})
}

// @Summary      Show the course
// @Param		 id path uint true "ID курса"
// @Failure		 400 {object} apperrors.ValidationError
// @Failure		 500 {object} map[string]string
// @Router		 /courses/{id} [get]
// @Tags 		 Courses
// @Success 	 200 {object} map[string]string
func (c2 *CourseHandler) GetCourse(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		_ = c.Error(apperrors.NewValidationError("некорректный ID курса"))
		c.Abort()
		return
	}
	logrus.Info("Получение курса")
	course, err := c2.courseService.GetCourse(uint(id))
	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, course)
}

// @Summary      Get all courses
// @Failure		 500 {object} map[string]string
// @Router		 /courses [get]
// @Tags 		 Courses
// @Success 	 200 {array} dtos.CourseDTO
func (c2 *CourseHandler) GetAllCourses(c *gin.Context) {
	logrus.Info("Получение списка всех курсов")
	courses, err := c2.courseService.GetAllCourses()
	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, courses)
}

// @Summary      Course update
// @Param		 input body dtos.CourseDTO true "Новые данные курса"
// @Failure		 400 {object} apperrors.ValidationError
// @Failure		 500 {object} map[string]string
// @Router		 /courses/{id} [put]
// @Tags 		 Courses
// @Success 	 200 {object} map[string]string
func (c2 *CourseHandler) UpdateCourse(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		_ = c.Error(apperrors.NewValidationError("некорректный ID курса"))
		c.Abort()
		return
	}
	var course dtos.CourseDTO
	if err = c.ShouldBindJSON(&course); err != nil {
		_ = c.Error(apperrors.NewValidationError("некорректный формат запроса"))
		c.Abort()
		return
	}
	logrus.Info("Обновление текущего курса")
	course.Id = uint(id)
	err = c2.courseService.UpdateCourse(course)
	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, course)
}

// @Summary      Delete a course
// @Param		 id path uint true "ID курса"
// @Failure		 400 {object} apperrors.ValidationError
// @Failure		 500 {object} map[string]string
// @Router		 /courses/{id} [delete]
// @Tags 		 Courses
// @Success 	 200 {object} map[string]string
func (c2 *CourseHandler) DeleteCourse(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		_ = c.Error(apperrors.NewValidationError("некорректный ID курса"))
		c.Abort()
		return
	}
	logrus.Info("Удаление текущего курса")
	err = c2.courseService.DeleteCourse(uint(id))
	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "курс успешно удален"})
}
