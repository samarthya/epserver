# Docker image.
FROM golang:1.14.7-alpine AS stage1
WORKDIR /go/src

COPY /src/* ./
COPY /src/msg/* ./msg/

RUN go build -o main server.go

# Stage 2
FROM alpine:latest

WORKDIR /service

# Copy the executable
COPY --from=stage1 /go/src/main .

# Server listens on 8090
EXPOSE 8090

CMD ["./main"]
