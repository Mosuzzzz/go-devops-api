# Stage 1: Build binary
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

# Stage 2: Final image (เล็กมาก)
FROM alpine:latest
WORKDIR /root/
# ก๊อปปี้มาเฉพาะไฟล์ binary ที่ build เสร็จแล้ว
COPY --from=builder /app/main . 
EXPOSE 8080
CMD ["./main"]