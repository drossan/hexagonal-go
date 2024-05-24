# Usa la imagen oficial de Golang como imagen base
FROM golang:1.22

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /go/src/github.com/drossan/core-api

# Copia los archivos del proyecto al directorio de trabajo del contenedor
COPY . .

# Descarga las dependencias
RUN go mod tidy

# Compila la aplicación
RUN go build -o hexagonal-go ./cmd/server

# Exponer el puerto 8080
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./hexagonal-go"]
