FROM golang:1.23.5

# Устанавливаем рабочую директорию
WORKDIR /avito-shop

# Копируем файлы зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Переходим в директорию с main.go и собираем приложение
RUN cd cmd/service && go build -o avito-shop .

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./cmd/service/avito-shop"]