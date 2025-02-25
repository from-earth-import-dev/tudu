FROM golang:1.19-alpine

WORKDIR /app

# First copy just the go.mod and go.sum files (if they exist)
COPY go.mod go.sum* ./

# If go.mod doesn't exist, initialize a new module
RUN if [ ! -f go.mod ]; then \
        go mod init task-service; \
    else \
        go mod download; \
    fi

# Then copy the rest of the application
COPY . .

# Build the application
RUN go build -o main .

EXPOSE 8080

CMD ["./main"] 