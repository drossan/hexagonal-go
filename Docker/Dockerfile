# Usa la imagen oficial de MySQL
FROM mariadb:latest

# Establece las variables de entorno para la base de datos
ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=hexagonal_go
ENV MYSQL_USER=hexagonal_user
ENV MYSQL_PASSWORD=password

# Copia el script de inicialización de la base de datos en el contenedor
COPY ./init.sql /docker-entrypoint-initdb.d/

# Exponer el puerto 3306
EXPOSE 3306