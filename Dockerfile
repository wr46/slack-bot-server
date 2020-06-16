# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Fixed image
FROM golang:1.13.1-alpine3.10

# Build Args
ARG APP_NAME=slack-bot-server
ARG LOG_DIR=/${APP_NAME}/logs

# Create Log Directory
RUN mkdir -p ${LOG_DIR}

# Environment Variables
ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log 

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Declare volumes to mount
VOLUME ["/slack-bot-server/logs"]

EXPOSE 3000

# Command to run the executable
CMD ["./main"]