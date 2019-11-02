FROM golang:1.13-alpine AS build

# Metadata for the build, compile only for linux target, smaller binary size
ENV GOOS linux
ENV GOARCH amd64
# Turn off CGO since that can result in dynamic links to libc/libmusl.
ENV CGO_ENABLED 0 

WORKDIR /app 

# Copy `go.mod` for definitions and `go.sum` to invalidate the next layer
# in case of a change in the dependencies
COPY src/go.mod src/go.sum ./
# Download dependencies
RUN go mod download

# Copy and build the app
COPY src/ .
CMD ["go", "run", "main.go"] 