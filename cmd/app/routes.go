package main

import (
	"lms-main-service/internal/handlers"
	"lms-main-service/pkg/apperrors"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	courseHandler *handlers.CourseHandler,
	chapterHandler *handlers.ChapterHandler,
	lessonHandler *handlers.LessonHandler,
) *gin.Engine {
	r := gin.Default()
	r.Use(apperrors.CustomErrMiddleware())

	api := r.Group("/api/v1")
	{
		// Курсы
		api.POST("/courses", courseHandler.CreateCourse)
		api.GET("/courses", courseHandler.GetAllCourses)
		api.GET("/courses/:id", courseHandler.GetCourse)
		api.PUT("/courses/:id", courseHandler.UpdateCourse)
		api.DELETE("/courses/:id", courseHandler.DeleteCourse)

		// Главы
		api.POST("/chapters", chapterHandler.CreateChapter)
		api.GET("/chapters/:id", chapterHandler.GetChapter)
		api.GET("/courses/:id/chapters", chapterHandler.GetChaptersByCourseID)
		api.PUT("/chapters/:id", chapterHandler.UpdateChapter)
		api.DELETE("/chapters/:id", chapterHandler.DeleteChapter)

		// Уроки
		api.POST("/lessons", lessonHandler.CreateLesson)
		api.GET("/lessons/:id", lessonHandler.GetLesson)
		api.GET("/chapters/:id/lessons", lessonHandler.GetLessonsByChapterID)
		api.PUT("/lessons/:id", lessonHandler.UpdateLesson)
		api.DELETE("/lessons/:id", lessonHandler.DeleteLesson)
	}
	return r
}
