# Build Stage
FROM golang:1.22.4-alpine3.20 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the application
RUN go build -o server .

# Run Stage
FROM alpine:3.20

WORKDIR /app

COPY --from=build /app/server .

EXPOSE 7000

CMD ["./server"]
