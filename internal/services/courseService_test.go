package services

import (
	"lms-main-service/internal/dtos"
	"lms-main-service/internal/mocks"
	"lms-main-service/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateCourse(t *testing.T) {
	// 1. Создаём структуру DTO с тестовыми данными
	testDTO := dtos.CourseDTO{
		Name:        "Go для начинающих",
		Description: "Курс от Bitlab",
	}

	// 2. Создаем поддельный репозиторий
	mockRepo := new(mocks.CourseRepository)

	// 3. Создаем НАСТОЯЩИЙ сервис, но подсовываем ему подделку
	courseService := NewCourseService(mockRepo)
	//Говорим ему: ты должен вернуть (.Return) пустой курс nil и ошибку, что запись не найдена
	mockRepo.On("GetBy", "name = ?", testDTO.Name).Return(nil, gorm.ErrRecordNotFound)
	// Говорим ему: "Когда сервис вызовет метод CreateCourse с любыми данными, верни nil (ошибок нет)"
	mockRepo.On("CreateCourse", mock.Anything).Return(nil)

	err := courseService.CreateCourse(testDTO)
	// 6. Проверяем, что метод сервиса не вернул ошибку
	assert.NoError(t, err)

	// 7. Проверяем, что сервис РЕАЛЬНО вызвал метод мока, который мы ожидали
	mockRepo.AssertExpectations(t)
}

func TestGetCourse(t *testing.T) {
	testID := uint(1)
	MockCourse := &models.Course{
		Id:          testID,
		Name:        "Go для начинающих",
		Description: "какое-то описание",
	}
	expectedDTO := dtos.CourseDTO{
		Id:          testID,
		Name:        "Go для начинающих",
		Description: "какое-то описание",
	}

	mockRepo := new(mocks.CourseRepository)

	courseService := NewCourseService(mockRepo)
	mockRepo.On("GetCourse", testID).Return(MockCourse, nil)
	courseDTO, err := courseService.GetCourse(testID)

	assert.NoError(t, err)
	assert.Equal(t, expectedDTO, courseDTO)
	mockRepo.AssertExpectations(t)
}

func TestGetAllCourses(t *testing.T) {
	CoursesDTO := []dtos.CourseDTO{}
	MockCourses := &[]models.Course{}
	mockRepo := new(mocks.CourseRepository)
	courseService := NewCourseService(mockRepo)
	mockRepo.On("GetAllCourses").Return(MockCourses, nil)
	courses, err := courseService.GetAllCourses()
	assert.NoError(t, err)
	assert.Len(t, courses, len(CoursesDTO))
	mockRepo.AssertExpectations(t)
}

func TestModifyCourse(t *testing.T) {
	testID := uint(1)
	mockCourse := &models.Course{
		Id:          testID,
		Name:        "старое название",
		Description: "старое описание",
	}
	updateCourse := dtos.CourseDTO{
		Id:          testID,
		Name:        "Go для начинающих",
		Description: "какое-то описание",
	}

	mockRepo := new(mocks.CourseRepository)
	courseService := NewCourseService(mockRepo)
	mockRepo.On("GetBy", "id = ?", testID).Return(mockCourse, nil)
	mockRepo.On("UpdateCourse", mock.Anything).Return(nil)

	err := courseService.UpdateCourse(updateCourse)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteCourse(t *testing.T) {
	testID := uint(1)
	mockCourse := &models.Course{
		Id:          testID,
		Name:        "старое название",
		Description: "старое описание",
	}
	mockRepo := new(mocks.CourseRepository)
	courseService := NewCourseService(mockRepo)
	mockRepo.On("GetBy", "id = ?", testID).Return(mockCourse, nil)
	mockRepo.On("DeleteCourse", testID).Return(nil)
	err := courseService.DeleteCourse(testID)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
