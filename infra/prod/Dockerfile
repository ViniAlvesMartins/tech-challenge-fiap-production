FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

EXPOSE 8080

RUN go build -o ./api/main ./api/main.go

CMD [ "./api/main" ]