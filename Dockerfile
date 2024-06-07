# Build Stage
FROM golang:1.22.4-alpine3.20 AS build

WORKDIR /app

COPY . .

# Download dependencies
RUN go get -d -v ./...

# Build the application
RUN go build -o server .

# Run Stage
FROM alpine:3.20

WORKDIR /app

COPY --from=build /app/server .

EXPOSE 7000

CMD ["./server"]
