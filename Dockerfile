# Stage 1 -> Installer && go build
FROM golang:1.24.3 as builder

# Instala herramientas necesarias
RUN apk update && apk add --no-cache alpine-sdk git

# Directorio de trabajo
WORKDIR /app

# Copia archivos de dependencias y descarga módulos
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copia el resto del código
COPY . .

# Compila el binario, lo llamamos "app"
RUN go build -o app .

# Stage 2 -> Run
FROM alpine:latest

# (opcional) Instala librerías si las necesitas
RUN apk update && rm -rf /var/cache/apk/*

WORKDIR /app

# Copia el binario desde la etapa anterior
COPY --from=builder /app/app .

# Puerto expuesto por tu app
EXPOSE 8100

# Comando de entrada
ENTRYPOINT ["./app"]
