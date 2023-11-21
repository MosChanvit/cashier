# Use the official Golang image as the base image
FROM golang:1.21.3 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies efficiently
# COPY go.mod ./

COPY . ./

# Download Go module dependencies
RUN go mod tidy

RUN go mod download

# Copy the rest of the application source code to the container

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata libc6-compat

COPY --from=builder /app/main .

COPY --from=builder /app/config.yaml .
# RUN chmod +x /app/cmd/main




EXPOSE 80

CMD ["./main"]