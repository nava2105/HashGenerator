# Usar imagen base de Go para la compilación
FROM golang:1.23 AS builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el archivo Go al contenedor
COPY main.go .

# Compilar el ejecutable en modo estático para una imagen mínima posterior
RUN CGO_ENABLED=0 GOOS=linux go build -o hashgenerator main.go

# Crear la imagen final
FROM alpine:latest

# Instalar dependencias necesarias en la imagen mínima
RUN apk --no-cache add ca-certificates

# Copiar el ejecutable desde la imagen builder
COPY --from=builder /app/hashgenerator /hashgenerator

# Exponer el puerto de la aplicación
EXPOSE 8080

# Ejecutar la aplicación
CMD ["/hashgenerator"]
