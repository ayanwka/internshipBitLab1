package services

import (
	"lms-main-service/internal/dtos"
	"lms-main-service/internal/mocks"
	"lms-main-service/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateLesson(t *testing.T) {
	mockLesson := &models.Lesson{
		Name:        "базовый цикл for",
		Description: "какое-то описание",
		Content:     "какое-то содержимое",
		ChapterId:   1,
		LessonOrder: 3,
	}
	lessonDTO := dtos.LessonDTO{
		Name:        "базовый цикл for",
		Description: "какое-то описание",
		Content:     "какое-то содержимое",
		ChapterID:   1,
		LessonOrder: 3,
	}
	mockRepo := new(mocks.LessonRepository)
	service := NewLessonService(mockRepo)
	mockRepo.On("CreateLesson", mockLesson).Return(nil)
	err := service.CreateLesson(lessonDTO)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetLesson(t *testing.T) {
	testID := uint(1)

	mockLesson := &models.Lesson{
		Name:        "базовый цикл for",
		Description: "какое-то описание",
		Content:     "какое-то содержимое",
		ChapterId:   1,
		LessonOrder: 3,
	}
	lessonDTO := dtos.LessonDTO{
		Name:        "базовый цикл for",
		Description: "какое-то описание",
		Content:     "какое-то содержимое",
		ChapterID:   1,
		LessonOrder: 3,
	}
	mockRepo := new(mocks.LessonRepository)
	service := NewLessonService(mockRepo)
	mockRepo.On("GetLesson", testID).Return(mockLesson, nil)
	lesson, err := service.GetLesson(testID)

	assert.NoError(t, err)
	assert.Equal(t, lessonDTO, lesson)
	mockRepo.AssertExpectations(t)
}

func TestGetLessonsByChapterID(t *testing.T) {
	testID := uint(1)

	mockLesson := &models.Lesson{}
	mockLessonList := &[]models.Lesson{}

	mockRepo := new(mocks.LessonRepository)
	service := NewLessonService(mockRepo)
	mockRepo.On("GetBy", "id = ?", testID).Return(mockLesson, nil)
	mockRepo.On("GetLessonsByChapterID", testID).Return(mockLessonList, nil)

	lessons, err := service.GetLessonsByChapterID(testID)
	assert.NoError(t, err)
	assert.Empty(t, lessons)
	mockRepo.AssertExpectations(t)
}

func TestModifyLesson(t *testing.T) {
	testId := uint(1)

	mockLesson := &models.Lesson{
		Id:          testId,
		Name:        "базовый цикл for",
		Description: "старое описание",
		Content:     "старое содержимое",
		ChapterId:   1,
		LessonOrder: 3,
	}
	updLesson := dtos.LessonDTO{
		Id:          testId,
		Name:        "базовый цикл for",
		Description: "научимся использовать цикл for",
		Content:     "for i:=0",
		ChapterID:   1,
		LessonOrder: 3,
	}

	mockRepo := new(mocks.LessonRepository)
	service := NewLessonService(mockRepo)
	mockRepo.On("GetBy", "id = ?", testId).Return(mockLesson, nil)
	mockRepo.On("UpdateLesson", mock.Anything).Return(nil)

	err := service.UpdateLesson(updLesson)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteLesson(t *testing.T) {
	testID := uint(1)

	mockLesson := &models.Lesson{
		Id:          testID,
		Name:        "базовый цикл for",
		Description: "старое описание",
		Content:     "старое содержимое",
		ChapterId:   1,
		LessonOrder: 3,
	}

	mockRepo := new(mocks.LessonRepository)
	service := NewLessonService(mockRepo)
	mockRepo.On("GetBy", "id = ?", testID).Return(mockLesson, nil)
	mockRepo.On("DeleteLesson", testID).Return(nil)

	err := service.DeleteLesson(testID)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
