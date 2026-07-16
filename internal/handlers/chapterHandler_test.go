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

func Init() {
	gin.SetMode(gin.TestMode)
}

func TestCreateChapter(t *testing.T) {
	inputDTO := dtos.ChapterDTO{
		Name:         "Циклы в GO",
		Description:  "какое-то описание",
		CourseID:     1,
		ChapterOrder: 5,
	}
	jsonBody, err := json.Marshal(inputDTO)
	assert.NoError(t, err)

	mockService := new(mocks.ChapterService)
	mockService.On("CreateChapter", mock.Anything).Return(nil)

	handler := NewChapterHandler(mockService)
	router := gin.New()
	router.POST("/chapters", handler.CreateChapter)

	req, err := http.NewRequest(http.MethodPost, "/chapters", bytes.NewBuffer(jsonBody))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	var responseBody map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "глава успешно создана", responseBody["message"])
	mockService.AssertExpectations(t)
}

func TestGetChapter(t *testing.T) {
	testID := uint(1)

	chapterDTO := dtos.ChapterDTO{
		Id:           testID,
		Name:         "Циклы в GO",
		Description:  "какое-то описание",
		CourseID:     1,
		ChapterOrder: 5,
	}

	mockService := new(mocks.ChapterService)
	mockService.On("GetChapter", testID).Return(chapterDTO, nil)

	handler := NewChapterHandler(mockService)
	router := gin.New()
	router.GET("/chapters/:id", handler.GetChapter)

	req, err := http.NewRequest(http.MethodGet, "/chapters/1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var actualChapter dtos.ChapterDTO
	err = json.Unmarshal(w.Body.Bytes(), &actualChapter)
	assert.NoError(t, err)
	assert.Equal(t, chapterDTO, actualChapter)
	mockService.AssertExpectations(t)
}

func TestGetChaptersByCourseID(t *testing.T) {
	testID := uint(1)

	chapterDTOs := []dtos.ChapterDTO{}

	mockService := new(mocks.ChapterService)
	mockService.On("GetChaptersByCourseID", testID).Return(chapterDTOs, nil)

	handler := NewChapterHandler(mockService)
	router := gin.New()
	router.GET("/courses/:id/chapters", handler.GetChaptersByCourseID)

	req, err := http.NewRequest(http.MethodGet, "/courses/1/chapters", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var actualChapterList []dtos.ChapterDTO
	err = json.Unmarshal(w.Body.Bytes(), &actualChapterList)
	assert.NoError(t, err)
	assert.Empty(t, actualChapterList)
	mockService.AssertExpectations(t)
}

func TestModifyChapter(t *testing.T) {
	testID := uint(1)

	updChapter := dtos.ChapterDTO{
		Id:           testID,
		Name:         "Циклы в GO",
		Description:  "научитесь использовать цикл for",
		CourseID:     1,
		ChapterOrder: 5,
	}
	inputDTO, err := json.Marshal(updChapter)
	assert.NoError(t, err)

	mockService := new(mocks.ChapterService)
	mockService.On("UpdateChapter", mock.Anything).Return(nil)

	handler := NewChapterHandler(mockService)
	router := gin.New()
	router.PUT("/chapters/:id", handler.UpdateChapter)

	req, err := http.NewRequest(http.MethodPut, "/chapters/1", bytes.NewBuffer(inputDTO))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// ДОПОЛНИТЕЛЬНАЯ ПРОВЕРКА:
	// Декодируем то, что вернул хендлер в теле ответа
	var actualChapter dtos.ChapterDTO
	err = json.Unmarshal(w.Body.Bytes(), &actualChapter)
	assert.NoError(t, err)

	// Проверяем, что обновленные данные вернулись без искажений
	assert.Equal(t, updChapter, actualChapter)

	mockService.AssertExpectations(t)
}

func TestDeleteChapter(t *testing.T) {
	testID := uint(1)

	mockService := new(mocks.ChapterService)
	mockService.On("DeleteChapter", testID).Return(nil)

	handler := NewChapterHandler(mockService)
	router := gin.New()
	router.DELETE("/chapters/:id", handler.DeleteChapter)

	req, err := http.NewRequest(http.MethodDelete, "/chapters/1", bytes.NewBuffer(nil))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}
