FROM golang:1.24.6-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gin-app main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/gin-app .

EXPOSE 8080

CMD ["./gin-app"]
