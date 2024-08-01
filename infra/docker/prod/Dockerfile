FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

EXPOSE 8082

RUN go build -o ./cmd/api/main ./cmd/api/main.go

CMD [ "./cmd/api/main" ]