# Stage 1: Build binary
FROM golang:1.25-alpine AS builder
WORKDIR /app
# ก๊อปปี้ go.mod/go.sum ก่อน เพื่อให้ Docker cache layer ของ dependencies ได้
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o main .

# Stage 2: Final image (เล็กมาก)
FROM alpine:latest
WORKDIR /app
# ก๊อปปี้มาเฉพาะไฟล์ binary ที่ build เสร็จแล้ว
COPY --from=builder /app/main .
# รันด้วย user ธรรมดา ไม่ใช่ root เพื่อความปลอดภัย
RUN adduser -D appuser
USER appuser
EXPOSE 8080
CMD ["./main"]
