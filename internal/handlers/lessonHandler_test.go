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

func TestCreateLesson(t *testing.T) {
	inputDTO := dtos.LessonDTO{
		Name:        "базовый цикл for",
		Description: "какое-то описание",
		Content:     "какое-то содержимое",
		ChapterID:   1,
		LessonOrder: 3,
	}
	jsonBody, err := json.Marshal(inputDTO)
	assert.NoError(t, err)

	mockService := new(mocks.LessonService)
	mockService.On("CreateLesson", mock.Anything).Return(nil)

	handler := NewLessonHandler(mockService)
	router := gin.New()
	router.POST("/lessons", handler.CreateLesson)

	req, err := http.NewRequest(http.MethodPost, "/lessons", bytes.NewBuffer(jsonBody))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	var responseBody map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "лекция успешно создана", responseBody["message"])
	mockService.AssertExpectations(t)
}

func TestGetLesson(t *testing.T) {
	testID := uint(1)

	lessonDTO := dtos.LessonDTO{
		Id:          testID,
		Name:        "базовый цикл for",
		Description: "какое-то описание",
		Content:     "какое-то содержимое",
		ChapterID:   1,
		LessonOrder: 3,
	}

	mockService := new(mocks.LessonService)
	mockService.On("GetLesson", testID).Return(lessonDTO, nil)

	handler := NewLessonHandler(mockService)
	router := gin.New()
	router.GET("/lessons/:id", handler.GetLesson)

	req, err := http.NewRequest(http.MethodGet, "/lessons/1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var actualLesson dtos.LessonDTO
	err = json.Unmarshal(w.Body.Bytes(), &actualLesson)
	assert.NoError(t, err)
	assert.Equal(t, lessonDTO, actualLesson)
	mockService.AssertExpectations(t)
}

func TestGetLessonsByChapterID(t *testing.T) {
	testID := uint(1)

	lessonDTOs := []dtos.LessonDTO{}

	mockService := new(mocks.LessonService)
	mockService.On("GetLessonsByChapterID", testID).Return(lessonDTOs, nil)

	handler := NewLessonHandler(mockService)
	router := gin.New()
	router.GET("/chapters/:id/lessons", handler.GetLessonsByChapterID)

	req, err := http.NewRequest(http.MethodGet, "/chapters/1/lessons", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var actualLessonList []dtos.LessonDTO
	err = json.Unmarshal(w.Body.Bytes(), &actualLessonList)
	assert.NoError(t, err)
	assert.Empty(t, actualLessonList)
	mockService.AssertExpectations(t)
}

func TestModifyLesson(t *testing.T) {
	testID := uint(1)

	updLesson := dtos.LessonDTO{
		Id:          testID,
		Name:        "базовый цикл for",
		Description: "научимся использовать цикл for",
		Content:     "for i:=0",
		ChapterID:   1,
		LessonOrder: 3,
	}
	inputDTO, err := json.Marshal(updLesson)
	assert.NoError(t, err)

	mockService := new(mocks.LessonService)
	mockService.On("UpdateLesson", mock.Anything).Return(nil)

	handler := NewLessonHandler(mockService)
	router := gin.New()
	router.PUT("/lessons/:id", handler.UpdateLesson)

	req, err := http.NewRequest(http.MethodPut, "/lessons/1", bytes.NewBuffer(inputDTO))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Декодируем и проверяем возвращаемый JSON-объект
	var actualLesson dtos.LessonDTO
	err = json.Unmarshal(w.Body.Bytes(), &actualLesson)
	assert.NoError(t, err)

	assert.Equal(t, updLesson, actualLesson)
	mockService.AssertExpectations(t)
}

func TestDeleteLesson(t *testing.T) {
	testID := uint(1)

	mockService := new(mocks.LessonService)
	mockService.On("DeleteLesson", testID).Return(nil)

	handler := NewLessonHandler(mockService)
	router := gin.New()
	router.DELETE("/lessons/:id", handler.DeleteLesson)

	req, err := http.NewRequest(http.MethodDelete, "/lessons/1", bytes.NewBuffer(nil))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}
