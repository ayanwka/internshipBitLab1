package main

import (
	"lms-main-service/internal/handlers"
	"lms-main-service/internal/repositories"
	"lms-main-service/internal/services"
	"lms-main-service/pkg/logging"
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logging.Init()
	// Проверяем: если запущено в Докере, подставится переменная из docker-compose
	dsn := os.Getenv("DB_DSN") //Data Source Name (Имя источника данных) - данные для подключения к БД
	if dsn == "" {
		// Дефолтная строка для локального запуска без докера (порт 5433, так как в докере у тебя 5433)
		dsn = "host=localhost user=user password=password dbname=lms_db port=5433 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Panicf("Не удалось подключиться к базе данных: %v", err)
	}

	courseRepo := repositories.NewCourseRepository(db)
	courseService := services.NewCourseService(courseRepo)
	courseHandler := handlers.NewCourseHandler(courseService)

	chapterRepo := repositories.NewChapterRepository(db)
	chapterService := services.NewChapterService(chapterRepo)
	chapterHandler := handlers.NewChapterHandler(chapterService)

	lessonRepo := repositories.NewLessonRepository(db)
	lessonService := services.NewLessonService(lessonRepo)
	lessonHandler := handlers.NewLessonHandler(lessonService)

	r := SetupRouter(courseHandler, chapterHandler, lessonHandler)
	log.Println("Сервер успешно запущен на http://localhost:8080")
	if err = r.Run(":8080"); err != nil {
		logrus.Fatalf("не удалось запустить сервер: %v", err)
	}
}
