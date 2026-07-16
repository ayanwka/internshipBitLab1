package services

import (
	"lms-main-service/internal/dtos"
	"lms-main-service/internal/mocks"
	"lms-main-service/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// 1. Создаём структуру DTO с тестовыми данными
// 2. Создаем поддельный репозиторий
// 3. Создаем НАСТОЯЩИЙ сервис, но подсовываем ему подделку
// Когда сервис вызовет GetBy ты должен вернуть (.Return) пустой курс nil и ошибку, что запись не найдена
// "Когда сервис вызовет метод X с любыми данными, верни nil (ошибок нет)"
// 6. Проверяем, что метод сервиса не вернул ошибку
// 7. Проверяем, что сервис РЕАЛЬНО вызвал метод мока, который мы ожидали

func TestCreateChapter(t *testing.T) {
	testDTO := dtos.ChapterDTO{
		Name:         "Циклы в GO",
		Description:  "какое-то описание",
		CourseID:     1,
		ChapterOrder: 5,
	}
	mockRepo := new(mocks.ChapterRepository)
	service := NewChapterService(mockRepo)

	mockRepo.On("CreateChapter", mock.Anything).Return(nil)
	err := service.CreateChapter(testDTO)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetChapter(t *testing.T) {
	testID := uint(1)
	mockChapter := &models.Chapter{
		Id:           testID,
		Name:         "Циклы в GO",
		Description:  "какое-то описание",
		CourseId:     1,
		ChapterOrder: 5,
	}
	chapterDTO := dtos.ChapterDTO{
		Id:           testID,
		Name:         "Циклы в GO",
		Description:  "какое-то описание",
		CourseID:     1,
		ChapterOrder: 5,
	}
	mockRepo := new(mocks.ChapterRepository)
	chapterService := NewChapterService(mockRepo)

	mockRepo.On("GetChapter", testID).Return(mockChapter, nil)
	chapter, err := chapterService.GetChapter(testID)
	assert.NoError(t, err)
	assert.Equal(t, chapterDTO, chapter)
	mockRepo.AssertExpectations(t)
}

func TestGetChaptersByCourseID(t *testing.T) {
	testID := uint(1)

	mockChapter := &models.Chapter{}
	mockChapterList := &[]models.Chapter{}

	mockRepo := new(mocks.ChapterRepository)
	service := NewChapterService(mockRepo)

	mockRepo.On("GetBy", "id = ?", testID).Return(mockChapter, nil)
	mockRepo.On("GetChaptersByCourseID", testID).Return(mockChapterList, nil)

	chapter, err := service.GetChaptersByCourseID(testID)
	assert.NoError(t, err)
	assert.Empty(t, chapter)
	mockRepo.AssertExpectations(t)
}

func TestModifyChapter(t *testing.T) { //название Modify, а не Update, ведь винда ругается
	testID := uint(1)

	// Имитируем существующую в БД модель главы (до обновления)
	mockChapter := &models.Chapter{
		Id:           testID,
		Name:         "старое имя",
		Description:  "старое описание",
		CourseId:     1,
		ChapterOrder: 5,
	}

	// Входное DTO с новыми данными, которые прислал клиент для обновления
	updChapter := dtos.ChapterDTO{
		Id:           testID,
		Name:         "Циклы в GO",
		Description:  "глава для начинающих",
		CourseID:     1,
		ChapterOrder: 5,
	}

	mockRepo := new(mocks.ChapterRepository) // Создаем мок-репозиторий (заглушку БД)
	service := NewChapterService(mockRepo)   // Инициализируем тестируемый сервис

	// Обучаем мок: при проверке существования главы ("GetBy") возвращаем нашу модель без ошибки
	mockRepo.On("GetBy", "id = ?", testID).Return(mockChapter, nil)

	// Обучаем мок: при вызове обновления принимаем любые параметры и возвращаем nil (успех)
	mockRepo.On("UpdateChapter", mock.Anything).Return(nil)

	// Вызываем реальный метод сервиса
	err := service.UpdateChapter(updChapter)

	// Проверяем, что метод завершился без ошибок
	assert.NoError(t, err)

	// Гарантируем, что сервис действительно вызвал все запрограммированные методы мока (GetBy и UpdateChapter)
	mockRepo.AssertExpectations(t)
}

func TestDeleteChapter(t *testing.T) {
	testID := uint(1)
	mockChapter := &models.Chapter{
		Id:           testID,
		Name:         "Циклы в GO",
		Description:  "какое-то описание",
		CourseId:     1,
		ChapterOrder: 5,
	}
	mockRepo := new(mocks.ChapterRepository)
	service := NewChapterService(mockRepo)
	mockRepo.On("GetBy", "id = ?", testID).Return(mockChapter, nil)
	mockRepo.On("DeleteChapter", testID).Return(nil)
	err := service.DeleteChapter(testID)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
