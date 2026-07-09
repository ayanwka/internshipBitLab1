# Этап 1: Сборка бинарника
FROM golang:1.25-alpine AS builder

# Устанавливаем рабочую папку внутри контейнера
WORKDIR /app

# Сначала копируем файлы зависимостей go.mod и go.sum
COPY go.mod go.sum ./

# Скачиваем библиотеки (это закешируется и сборка будет быстрой)
RUN go mod download

# Копируем весь остальной код проекта
COPY . .

# Собираем скомпилированный бинарный файл под именем "lms_app"
RUN CGO_ENABLED=0 GOOS=linux go build -o lms_app ./cmd/app

# Этап 2: Легковесный запуск
FROM alpine:latest

WORKDIR /app

# Копируем только готовый скомпилированный файл из предыдущего этапа
COPY --from=builder /app/lms_app .

# Говорим контейнеру запускать этот файл на старте
CMD ["./lms_app"]