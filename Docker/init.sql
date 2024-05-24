CREATE DATABASE IF NOT EXISTS hexagonal_go;

DROP USER IF EXISTS 'hexagonal_user'@'%';
CREATE USER 'hexagonal_user'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON hexagonal_go.* TO 'hexagonal_user'@'%';
FLUSH PRIVILEGES;