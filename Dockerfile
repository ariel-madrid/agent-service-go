# Stage 1 -> Installer && go build
FROM golang:1.24.3 as builder

RUN apt-get update && \
    apt-get install -y --no-install-recommends build-essential git && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o app .

# Stage 2 -> Run
FROM alpine:latest

RUN apk update && rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8100

ENTRYPOINT ["./app"]
