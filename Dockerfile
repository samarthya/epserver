# Docker image.
FROM golang:1.14.7-alpine AS stage1

WORKDIR /go/src

COPY /src/* ./
COPY /src/msg/* ./msg/

# Version 2 Added some HTML file handling
COPY /src/wiki/* ./wiki/
COPY /src/templates/* ./templates/

# Build the binary
RUN go build -o /build/main server.go

# Stage 2
FROM alpine:latest

WORKDIR /service

# Copy the executable
COPY --from=stage1 ./build/main .
COPY /src/templates/* ./templates/

# Server listens on 8090
EXPOSE 8090

ENTRYPOINT [ "./main" ]
