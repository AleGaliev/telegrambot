# Этап 1: Сборка приложения
FROM golang:1.22.2 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем исходный код в контейнер
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Этап 2: Создание финального образа на основе ubi
FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем бинарный файл из этапа сборки
COPY --from=builder /app/myapp .

# Указываем команду для запуска приложения
CMD ["./myapp"]