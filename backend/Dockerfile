FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o tourism-backend ./src/cmd/


FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /app/tourism-backend ./tourism-backend

EXPOSE 8080

CMD ["./tourism-backend"]