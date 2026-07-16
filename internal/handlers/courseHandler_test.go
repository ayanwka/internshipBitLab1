package handlers

import (
	"bytes"
	"encoding/json"
	"lms-main-service/internal/dtos"
	"lms-main-service/internal/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCourse(t *testing.T) {
	courseDTO := dtos.CourseDTO{
		Name:        "Основы языка GO",
		Description: "какое-то описание",
	}
	jsonInput, err := json.Marshal(courseDTO)
	assert.NoError(t, err)

	mockService := new(mocks.CourseService)
	mockService.On("CreateCourse", mock.Anything).Return(nil)

	handler := NewCourseHandler(mockService)
	router := gin.New()
	router.POST("/courses", handler.CreateCourse)

	req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(jsonInput))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetCourse(t *testing.T) {
	testID := uint(1)

	courseDTO := dtos.CourseDTO{
		Id:          testID,
		Name:        "Основы языка GO",
		Description: "какое-то описание",
	}

	mockService := new(mocks.CourseService)
	mockService.On("GetCourse", testID).Return(courseDTO, nil)

	handler := NewCourseHandler(mockService)
	router := gin.New()
	router.GET("/courses/:id", handler.GetCourse)

	req, err := http.NewRequest("GET", "/courses/1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var actualCourse dtos.CourseDTO
	err = json.Unmarshal(w.Body.Bytes(), &actualCourse)
	assert.NoError(t, err)
	assert.Equal(t, courseDTO, actualCourse)
	mockService.AssertExpectations(t)
}

func TestGetAllCourses(t *testing.T) {
	courseDTOs := []dtos.CourseDTO{}

	mockService := new(mocks.CourseService)
	mockService.On("GetAllCourses").Return(courseDTOs, nil)

	handler := NewCourseHandler(mockService)
	router := gin.New()
	router.GET("/courses", handler.GetAllCourses)

	req, err := http.NewRequest("GET", "/courses", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestModifyCourse(t *testing.T) {
	testID := uint(1)

	updCourse := dtos.CourseDTO{
		Id:          testID,
		Name:        "Функции в ГО",
		Description: "методы и функции",
	}
	jsonDTO, err := json.Marshal(updCourse)
	assert.NoError(t, err)

	mockService := new(mocks.CourseService)
	mockService.On("UpdateCourse", mock.Anything).Return(nil)

	handler := NewCourseHandler(mockService)
	router := gin.New()
	router.PUT("/courses/:id", handler.UpdateCourse)

	req, err := http.NewRequest("PUT", "/courses/1", bytes.NewBuffer(jsonDTO))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var actualCourse dtos.CourseDTO
	err = json.Unmarshal(w.Body.Bytes(), &actualCourse)
	assert.NoError(t, err)

	// Проверяем, что обновленные данные вернулись без искажений
	assert.Equal(t, updCourse, actualCourse)
	mockService.AssertExpectations(t)
}

func TestDeleteCourse(t *testing.T) {
	testID := uint(1)

	mockService := new(mocks.CourseService)
	mockService.On("DeleteCourse", testID).Return(nil)

	handler := NewCourseHandler(mockService)
	router := gin.New()
	router.DELETE("/courses/:id", handler.DeleteCourse)

	req, err := http.NewRequest(http.MethodDelete, "/courses/1", bytes.NewBuffer(nil))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}
