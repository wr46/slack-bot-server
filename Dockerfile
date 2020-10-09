FROM golang:1.15-alpine3.12 AS build

ENV CGO_ENABLED=0

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Allow to execute the script when App running
RUN ["chmod", "+x", "./scripts/script.sh"]

# Build the Go app
RUN go build -o main .

# Linter execution
FROM golangci/golangci-lint:v1.31-alpine AS linter

FROM build AS lint

COPY --from=linter /usr/bin/golangci-lint /usr/bin/golangci-lint

RUN golangci-lint run .

# Application binary execution
FROM build AS bin

COPY --from=build /app /app

EXPOSE 3000

# Command to run the executable
CMD ["./main"]