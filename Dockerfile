FROM golang:1.23.5

WORKDIR /avito-shop

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY --from=ghcr.io/ufoscout/docker-compose-wait:latest /wait /wait

WORKDIR /avito-shop/cmd/service

RUN go build -o avito-shop .

EXPOSE 8080

CMD /wait && ./avito-shop
