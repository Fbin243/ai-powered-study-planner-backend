FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN make build

RUN chmod +x main

EXPOSE 8080

CMD ["./main"]
