# Docker image.
FROM golang:1.14.7-alpine AS stage1

WORKDIR /go/src

COPY /src/* ./
COPY /src/msg/* ./msg/

# Build the binary
RUN go build -o /build/main server.go

# Stage 2
FROM alpine:latest

WORKDIR /service

# Copy the executable
COPY --from=stage1 ./build/main .

# Server listens on 8090
EXPOSE 8090

ENTRYPOINT [ "./main" ]
