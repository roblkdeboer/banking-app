# syntax=docker/dockerfile:1

FROM golang:1.20.6

WORKDIR /app

COPY go.mod ./
COPY *.go ./
COPY static ./static

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /banking-app

# Run
CMD ["/banking-app"]