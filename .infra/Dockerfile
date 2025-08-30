FROM  golang:1.22-alpine AS build

ARG VERSION=NONE
ARG ARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=$ARCH go build -o wb-semaphore ./cmd/wb-semaphore/main.go

FROM alpine:edge

WORKDIR /app

COPY --from=build /app/wb-semaphore wb-semaphore .