FROM golang:1.25-alpine AS builder

WORKDIR  /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o LavaProxy .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/LavaProxy .

EXPOSE 3001

CMD ["./LavaProxy"]