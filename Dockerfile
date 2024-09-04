# Указываем базовый образ с установленным Go
FROM golang:1.23-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /tms

# Копируем модульные файлы и загружаем зависимости
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Копируем остальной код приложения
COPY . .

# Собираем приложение
RUN go build -o main ./cmd/tms

RUN go build -o migrator ./cmd/migrator

# Определяем команду для запуска приложения
CMD ["./main"]