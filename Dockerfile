# syntax=docker/dockerfile:1

FROM golang:1.20.6

WORKDIR /banking-app

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
COPY static ./static
COPY postgres ./postgres

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /banking-app

# Run
CMD ["./banking-app"]