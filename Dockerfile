# syntax=docker/dockerfile:1

FROM golang:1.20.5 AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./main ./cmd

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /app

COPY --from=builder /usr/src/app/main .

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["./main"]