# Используем официальный образ Golang как базовый
FROM golang:1.23-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы модулей Go и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка приложения
RUN go build -o pflDockerBilling ./cmd/main.go

# Указываем порт, на котором будет работать приложение
EXPOSE 8080

# Запускаем приложение
CMD ["./pflDockerBilling"]