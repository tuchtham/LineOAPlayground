FROM golang:1.20.5-alpine

WORKDIR /app

COPY . .

EXPOSE 8080

RUN go mod tidy

CMD ["go", "run", "./main.go"]