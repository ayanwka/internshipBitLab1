package mappers

import (
	"lms-main-service/internal/dtos"
	"lms-main-service/internal/models"
)

func MapToCourseDto(course models.Course) dtos.CourseDTO {
	var courseDTO dtos.CourseDTO
	courseDTO.Id = course.Id
	courseDTO.Name = course.Name
	courseDTO.Description = course.Description
	return courseDTO
}

func MapToCourseDtoList(courses []models.Course) []dtos.CourseDTO {
	var courseDTOs []dtos.CourseDTO
	for i := 0; i < len(courses); i++ {
		var courseDTO dtos.CourseDTO
		courseDTO = MapToCourseDto(courses[i])
		courseDTOs = append(courseDTOs, courseDTO)
	}
	return courseDTOs
}

func MapToChapterDto(chapter models.Chapter) dtos.ChapterDTO {
	var chapterDTO dtos.ChapterDTO
	chapterDTO.Id = chapter.Id
	chapterDTO.Name = chapter.Name
	chapterDTO.Description = chapter.Description
	chapterDTO.CourseID = chapter.CourseId
	chapterDTO.ChapterOrder = chapter.ChapterOrder
	return chapterDTO
}

func MapToChapterDtoList(chapters []models.Chapter) []dtos.ChapterDTO {
	var chapterDTOs []dtos.ChapterDTO
	for i := 0; i < len(chapters); i++ {
		var chapterDTO dtos.ChapterDTO
		chapterDTO = MapToChapterDto(chapters[i])
		chapterDTOs = append(chapterDTOs, chapterDTO)
	}
	return chapterDTOs
}

func MapToLessonDto(lesson models.Lesson) dtos.LessonDTO {
	var lessonDTO dtos.LessonDTO
	lessonDTO.Id = lesson.Id
	lessonDTO.Name = lesson.Name
	lessonDTO.Description = lesson.Description
	lessonDTO.Content = lesson.Content
	lessonDTO.ChapterID = lesson.ChapterId
	lessonDTO.LessonOrder = lesson.LessonOrder
	return lessonDTO
}

func MapToDtoLessonList(lessons []models.Lesson) []dtos.LessonDTO {
	var lessonDTOs []dtos.LessonDTO
	for i := 0; i < len(lessons); i++ {
		var lessonDTO dtos.LessonDTO
		lessonDTO = MapToLessonDto(lessons[i])
		lessonDTOs = append(lessonDTOs, lessonDTO)
	}
	return lessonDTOs
}
