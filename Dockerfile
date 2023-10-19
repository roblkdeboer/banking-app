# syntax=docker/dockerfile:1

FROM golang:1.20.6

WORKDIR /banking-app

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
COPY static ./static
COPY handlers ./handlers
COPY models ./models
COPY postgres ./postgres
COPY utils ./utils
COPY users ./users

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /banking-app

EXPOSE 8080 443

# Run
ENTRYPOINT ["./banking-app"]