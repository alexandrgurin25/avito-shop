# Используем официальный образ Go
FROM golang:1.23.5

# Устанавливаем рабочую директорию
WORKDIR /avito-shop

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные файлы
COPY . .

COPY .env ./

WORKDIR /avito-shop/cmd/service

# Компилируем приложение
RUN go build -o avito-shop .

# Указываем, что контейнер слушает на порту 8080
EXPOSE 8080

# Запускаем приложение
CMD ["./avito-shop"]
