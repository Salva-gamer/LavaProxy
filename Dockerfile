FROM golang:1.25-alpine AS builder 


WORKDIR /app
COPY go.mod ./

RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o LavaProxy .

# --- Etapa Final ---
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/


COPY --from=builder /app/LavaProxy .


COPY --from=builder /app/config.yml .

EXPOSE 3001

CMD ["./LavaProxy"]