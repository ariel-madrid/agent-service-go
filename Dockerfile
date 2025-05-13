# Stage 1 -> Installer && go build
FROM golang:1.21.3-alpine as builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./app ./cmd/server/main.go

# Stage 2 -> Run
FROM alpine:latest

RUN apk update && rm -rf /var/cache/apk/*

RUN mkdir -p /app
WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8100

ENTRYPOINT [ "./app" ]