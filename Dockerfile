FROM golang:1.17.6-alpine AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "main" -ldflags="-w -s" ./main.go

FROM scratch

COPY --from=builder /app/main /usr/bin/

CMD ["main"]

EXPOSE 80