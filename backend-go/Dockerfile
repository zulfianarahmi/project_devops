# Build stage
FROM golang:1.20 AS build

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o main

# Run stage
FROM debian:bullseye-slim

WORKDIR /app
COPY --from=build /app/main .
EXPOSE 8080

CMD ["./main"]
