FROM golang:alpine as builder

ENV GO111MODULE=on
LABEL maintainer="Muhammad Talha <talhach891@gmail.com>"
# Install git, make
RUN apk update && apk add --no-cache git bash make

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN make

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY ./config-files/rest_task.ini  .

ENV REST_TASK_SETTINGS=/root/rest_task.ini

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/bin/rest .

EXPOSE 8081

LABEL Name=restserver Version=0.0.1

#Command to run the executable
CMD ["./rest", "serve_api"]
