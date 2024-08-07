FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

EXPOSE 8082

RUN go build -o ./cmd/api/main ./cmd/api/main.go
RUN go build -o ./cmd/worker/order_created_worker/main ./cmd/worker/order_created_worker/main.go
RUN go build -o ./cmd/worker/payment_confirmed_worker/main ./cmd/worker/payment_confirmed_worker/main.go

CMD [ "./cmd/api/main" ]