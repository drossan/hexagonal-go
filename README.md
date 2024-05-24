# Hexagonal Architecture Example API

Este proyecto es un ejemplo de una API construida utilizando la arquitectura hexagonal en Go, utilizando Echo y Gorm.

## Estructura del Proyecto

El proyecto está organizado en las siguientes capas:

- `cmd`: Contiene el punto de entrada de la aplicación.
- `config`: Carga y gestiona la configuración de la aplicación.
- `domain`: Contiene los modelos y repositorios de la aplicación.
- `infrastructure`: Contiene la implementación de los repositorios y la configuración de la base de datos.
- `interfaces`: Contiene los handlers de la API y la lógica de presentación.
- `usecase`: Contiene la lógica de negocio de la aplicación.
- `seeder`: Contiene los seeders para inicializar la base de datos con datos de ejemplo.
- `docs`: Contiene la documentación generada por Swagger.

## Configuración

1. **Instalar dependencias**:

    ```sh
    go mod tidy
    ```

2. **Configurar la base de datos**:

   Actualiza el archivo `config/config.yaml` con los detalles de tu base de datos.

3. **Inicializar la base de datos**:

    ```sh
    go run cmd/server/main.go
    ```

## Documentación de la API

La documentación de la API se genera utilizando Swagger. Para generar la documentación:

```sh
swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal
```

## Arquitectura Hexagonal
La arquitectura hexagonal promueve la separación de preocupaciones mediante la división del sistema en varias capas.
Esto permite que los componentes internos de la aplicación no dependan directamente de detalles de implementación como
frameworks web o bases de datos.

Estructura del Proyecto
```text
/hexagonal-example-go/
|-- .air.toml
|-- .env.example
|-- .github
|   |-- workflows
|   |   |-- go-tests.yml
|-- .gitignore
|-- Docker
|   |-- Dockerfile
|   |-- Dockerfile.app
|   |-- db_data
|   |-- docker-compose.yml
|   |-- init.sql
|-- README.md
|-- adapters
|   |-- email_notifier.go
|   |-- slack_notifier.go
|-- cmd
|   |-- server
|   |   |-- main.go
|-- config
|   |-- config.go
|-- docs
|   |-- docs.go
|   |-- swagger.json
|   |-- swagger.yaml
|-- domain
|   |-- model
|   |   |-- Clain.go
|   |   |-- LevelPrivileges.go
|   |   |-- form.go
|   |   |-- level.go
|   |   |-- menu_tree.go
|   |   |-- user.go
|   |-- notification
|   |   |-- notification.go
|   |   |-- notifier.go
|   |-- repository
|   |   |-- form_repository.go
|   |   |-- level_privileges_repository.go
|   |   |-- level_repository.go
|   |   |-- menu_tree_repository.go
|   |   |-- service_repository.go
|   |   |-- user_repository.go
|-- go.mod
|-- go.sum
|-- helpers
|   |-- jwt_helpers.go
|   |-- test
|   |   |-- jwt_helpers_test.go
|   |-- token.go
|-- http-requests
|   |-- auth.http
|   |-- forms.http
|   |-- hotels.http
|   |-- http-client.env.json
|   |-- levels.http
|   |-- users.http
|-- infrastructure
|   |-- db
|   |   |-- form_repository.go
|   |   |-- gorm.go
|   |   |-- level_privileges_repository.go
|   |   |-- level_repository.go
|   |   |-- menu_tree_repository.go
|   |   |-- migrations.go
|   |   |-- tests
|   |   |   |-- form_repository_test.go
|   |   |   |-- level_privileges_repository_test.go
|   |   |   |-- level_repository_test.go
|   |   |   |-- menu_tree_repository_test.go
|   |   |   |-- user_repository_test.go
|   |   |-- user_repository.go
|   |-- notification
|   |   |-- email_notification.go
|   |   |-- slack_notification.go
|   |-- router
|   |   |-- echo.go
|-- integration_tests
|   |-- integration_form_handler_test.go
|   |-- integration_level_handler_test.go
|   |-- integration_level_privileges_handler_test.go
|   |-- integration_menu_tree_handler_test.go
|   |-- integration_user_handler_test.go
|-- interfaces
|   |-- api
|   |   |-- form_handler.go
|   |   |-- level_handler.go
|   |   |-- level_privileges_handler.go
|   |   |-- menu_tree_handler.go
|   |   |-- tests
|   |   |   |-- form_handler_test.go
|   |   |   |-- level_handler_test.go
|   |   |   |-- level_privileges_handler_test.go
|   |   |   |-- menu_tree_handler_test.go
|   |   |   |-- mocks
|   |   |   |   |-- mock_user_repository.go
|   |   |   |-- user_handler_test.go
|   |   |-- user_handler.go
|-- middleware
|   |-- auth.go
|   |-- authorization.go
|   |-- rate_limiter.go
|-- mocks
|   |-- mock_form_repository.go
|   |-- mock_level_priviliges_repository.go
|   |-- mock_level_repository.go
|   |-- mock_menu_tree_repository.go
|   |-- mock_user_repository.go
|-- seeder
|   |-- seed.go
|-- service
|   |-- notification_service.go
|-- templates
|   |-- notification_email.html
|   |-- user_created.html
|-- usecase
|   |-- form_usecase.go
|   |-- level_privileges_usecase.go
|   |-- level_usecase.go
|   |-- menu_tree_usecase.go
|   |-- tests
|   |   |-- auth_usecase_test.go
|   |   |-- form_usecase_test.go
|   |   |-- level_privileges_usecase_test.go
|   |   |-- level_usecase_test.go
|   |   |-- menu_tree_usecase_test.go
|   |   |-- user_usecase_test.go
|   |-- user_usecase.go
|-- utils
|   |-- test_helpers.go
```

### Explicación de las Capas

1. cmd:

- Propósito: Contiene el punto de entrada de la aplicación.
- Detalle: main.go inicializa la aplicación, configura dependencias y arranca el servidor.

2. config:

- Propósito: Maneja la configuración de la aplicación.
- Detalle: config.go carga y proporciona la configuración necesaria para diferentes partes de la aplicación.

3. domain:

- Propósito: Contiene las entidades centrales y las interfaces de repositorio.
- Subcapas:
    - model: Define las estructuras de datos y las entidades del dominio. Ejemplos incluyen user.go y hotel.go.
    - repository: Define las interfaces que los repositorios deben implementar. Ejemplos incluyen user_repository.go y
      hotel_repository.go.

4. infrastructure:

- Propósito: Proporciona implementaciones concretas para las interfaces definidas en la capa de dominio.
- Subcapas:
    - db: Contiene implementaciones de repositorios que interactúan con la base de datos. Ejemplos incluyen
      gorm_user_repository.go y gorm_hotel_repository.go.
    - notification: Contiene implementaciones para servicios de notificación. Ejemplos incluyen email_notification.go y
      slack_notification.go.

5. interfaces:

- Propósito: Define adaptadores que exponen la funcionalidad de la aplicación a través de diferentes interfaces (por
  ejemplo, HTTP).
- Subcapas:
    - api: Contiene controladores que manejan las solicitudes HTTP y llaman a los casos de uso apropiados. Ejemplos
      incluyen user_handler.go y hotel_handler.go.
    - middleware: Define middleware para la aplicación. Ejemplos incluyen autenticación, logging, etc.

6. usecase:

- Propósito: Contiene la lógica de negocio central de la aplicación, orquestando las operaciones entre diferentes
  entidades y repositorios.
- Detalle: Define los casos de uso que encapsulan la lógica de negocio. Ejemplos incluyen user_usecase.go y
  hotel_usecase.go.

7. seeder:

- Propósito: Proporciona inicialización de datos para la aplicación.
- Detalle: seeder.go contiene lógica para poblar la base de datos con datos iniciales.

#### Detalle de las Capas con Ejemplos

1. cmd/server/main.go:

- Descripción: Punto de entrada principal de la aplicación. Configura y arranca el servidor.
- Ejemplo:

```go
package main

import (
	"github.com/drossan/hexagonal-core/config"
	"github.com/drossan/hexagonal-core/infrastructure/db"
	"github.com/drossan/hexagonal-core/interfaces/api"
	"github.com/drossan/hexagonal-core/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
	"os"
)

func main() {
	cfg := config.LoadConfig()
	dbConn := db.Initialize(cfg.Database.URL)
	db.Migrate(dbConn)

	// Crear servicios y repositorios
	userRepo := db.NewGormUserRepository(dbConn)
	userUseCase := usecase.NewUserUseCase(userRepo)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userHandler := api.NewUserHandler(e, userUseCase)
	userHandler.RegisterRoutes(e.Group("/api"))

	if os.Getenv("APP_ENV") != "production" {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	log.Fatal(e.Start(cfg.Server.Address))
}
```

2. domain/model/user.go:

- Descripción: Define la estructura de datos para un usuario.

```go
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
```

3. domain/repository/user_repository.go:

- Descripción: Define la interfaz para las operaciones de repositorio de usuario.

```go
package repository

import "github.com/drossan/hexagonal-core/domain/model"

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByID(id uint) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	DeleteUser(id uint) error
}
infrastructure
```

4. /db/gorm_user_repository.go:

- Descripción: Implementa la interfaz UserRepository usando GORM.

- Copiar código

```go
package db

import (
	"github.com/drossan/hexagonal-core/domain/model"
	"github.com/drossan/hexagonal-core/domain/repository"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repository.UserRepository {
	return &GormUserRepository{db}
}

func (r *GormUserRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *GormUserRepository) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) GetAllUsers() ([]*model.User, error) {
	var users []*model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *GormUserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}
```

5. interfaces/api/user_handler.go:

- Descripción: Maneja las solicitudes HTTP relacionadas con los usuarios.

```go
package usecase

import (
	"github.com/drossan/hexagonal-core/domain/model"
	"github.com/drossan/hexagonal-core/domain/repository"
	"github.com/drossan/hexagonal-core/domain/notification"
)

type UserUseCase struct {
	userRepository      repository.UserRepository
	notificationService *notification.NotificationService
}

func NewUserUseCase(repo repository.UserRepository, notificationService *notification.NotificationService) *UserUseCase {
	return &UserUseCase{
		userRepository:      repo,
		notificationService: notificationService,
	}
}

func (uc *UserUseCase) CreateUser(user *model.User) error {
	err := uc.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	// Enviar notificación por email y Slack
	message := "A new user has been created with email: " + user.Email
	notifierNames := []string{"email", "slack"}
	err = uc.notificationService.SendNotification(notifierNames, user.Email, message)
	if err != nil {
		return err
	}

	return nil
}

// Otros métodos del caso de uso...

```

6. config/config.go:

- Descripción: Maneja la carga y acceso a la configuración de la aplicación.

```go
package config

import (
	"log"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Email    EmailConfig
	Slack    SlackConfig
}

type ServerConfig struct {
	Address   string
	JWTSecret string
}

type DatabaseConfig struct {
	URL string
}

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPassword string
}

type SlackConfig struct {
	WebhookURL string
}

func LoadConfig() *Config {
	viper.SetConfigName("config") // Nombre del archivo de configuración (sin extensión)
	viper.SetConfigType("yaml")   // Tipo del archivo de configuración
	viper.AddConfigPath(".")      // Ruta donde se encuentra el archivo de configuración

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error leyendo archivo de configuración: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshaling configuración: %v", err)
	}

	return &config
}

```

7. infrastructure/notification/email_notification.go:

- Descripción: Implementa la interfaz Notification para enviar correos electrónicos.

```go
package notification

import (
	"fmt"
	"net/smtp"
)

type EmailNotification struct {
	smtpHost string
	smtpPort string
	auth     smtp.Auth
}

func NewEmailNotification(smtpHost, smtpPort, smtpUser, smtpPassword string) *EmailNotification {
	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)
	return &EmailNotification{
		smtpHost: smtpHost,
		smtpPort: smtpPort,
		auth:     auth,
	}
}

func (e *EmailNotification) Send(to string, message string) error {
	from := "youremail@example.com"
	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: Notification\n\n%s", from, to, message)
	err := smtp.SendMail(e.smtpHost+":"+e.smtpPort, e.auth, from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}

```

8. infrastructure/notification/slack_notification.go:

- Descripción: Implementa la interfaz Notification para enviar mensajes a Slack

```go
package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SlackNotification struct {
	webhookURL string
}

func NewSlackNotification(webhookURL string) *SlackNotification {
	return &SlackNotification{webhookURL: webhookURL}
}

func (s *SlackNotification) Send(to string, message string) error {
	payload := map[string]string{
		"channel": to,
		"text":    message,
	}
	payloadBytes, _ := json.Marshal(payload)

	resp, err := http.Post(s.webhookURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send slack notification, status code: %d", resp.StatusCode)
	}

	return nil
}

```

9. main.go:

- Descripción: Punto de entrada principal que configura y arranca el servidor.

```go
package main

import (
	"github.com/drossan/hexagonal-core/config"
	"github.com/drossan/hexagonal-core/infrastructure/db"
	"github.com/drossan/hexagonal-core/infrastructure/notification"
	"github.com/drossan/hexagonal-core/interfaces/api"
	"github.com/drossan/hexagonal-core/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
	"os"
)

func main() {
	// Cargar configuración
	cfg := config.LoadConfig()

	// Inicializar base de datos
	dbConn := db.Initialize(cfg.Database.URL)

	// Migrar base de datos
	db.Migrate(dbConn)

	// Crear el servicio de notificaciones
	notificationService := usecase.NewNotificationService()

	// Registrar notificador de email
	emailNotifier := notification.NewEmailNotification(cfg.Email.SMTPHost, cfg.Email.SMTPPort, cfg.Email.SMTPUser, cfg.Email.SMTPPassword)
	notificationService.RegisterNotifier("email", emailNotifier)

	// Registrar notificador de Slack
	slackNotifier := notification.NewSlackNotification(cfg.Slack.WebhookURL)
	notificationService.RegisterNotifier("slack", slackNotifier)

	// Crear instancias de los repositorios
	userRepo := db.NewGormUserRepository(dbConn)
	hotelRepo := db.NewGormHotelRepository(dbConn)
	menuTreeRepo := db.NewGormMenuTreeRepository(dbConn)

	// Inicializar casos de uso
	userUseCase := usecase.NewUserUseCase(userRepo, notificationService)
	hotelUseCase := usecase.NewHotelUseCase(hotelRepo)
	menuTreeUseCase := usecase.NewMenuTreeUseCase(menuTreeRepo)
	authUseCase := usecase.NewAuthUseCase(userRepo, cfg.Server.JWTSecret)

	// Inicializar router de Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Inicializar manejadores y registrar rutas
	authHandler := api.NewAuthHandler(e, authUseCase)
	userHandler := api.NewUserHandler(e, userUseCase)
	hotelHandler := api.NewHotelHandler(e, hotelUseCase, cfg.App.Env)
	menuTreeHandler := api.NewMenuTreeHandler(e, menuTreeUseCase)

	// Registrar rutas públicas
	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)

	// Configurar autenticación JWT para proteger rutas de la API
	jwtConfig := middleware.JWTConfig{
		SigningKey: []byte(cfg.Server.JWTSecret),
	}
	jwtMiddleware := middleware.JWTWithConfig(jwtConfig)

	// Proteger las rutas de la API
	apiGroup := e.Group("/api")
	apiGroup.Use(jwtMiddleware)
	userHandler.RegisterRoutes(apiGroup)
	hotelHandler.RegisterRoutes(apiGroup)
	menuTreeHandler.RegisterRoutes(apiGroup)

	// Habilitar Swagger solo en desarrollo
	if os.Getenv("APP_ENV") != "production" {
		swaggerGroup := e.Group("/swagger")
		swaggerGroup.Use(jwtMiddleware)
		swaggerGroup.GET("/*", echoSwagger.WrapHandler)
	}

	// Iniciar el servidor
	log.Fatal(e.Start(cfg.Server.Address))
}

```

##### Resumen

- cmd: Contiene el punto de entrada de la aplicación, configurando y arrancando el servidor.
- config: Maneja la carga y acceso a la configuración de la aplicación.
- domain: Contiene las entidades centrales y las interfaces de repositorio.
- infrastructure: Proporciona implementaciones concretas para las interfaces definidas en la capa de dominio (por
  ejemplo, repositorios y servicios de notificación).
- interfaces: Define adaptadores que exponen la funcionalidad de la aplicación a través de diferentes interfaces (por
  ejemplo, HTTP).
- usecase: Contiene la lógica de negocio central de la aplicación, orquestando las operaciones entre diferentes
  entidades y repositorios.
  Cada capa tiene una responsabilidad clara, lo que facilita la mantenibilidad y la escalabilidad del sistema. Esta
  organización permite que diferentes partes de la aplicación sean modificadas o extendidas sin afectar
  significativamente a otras partes, promoviendo un desarrollo más ágil y robusto.

## TEST

Los tests para los casos de uso (usecases), los repositorios (repositories), los handlers (controladores) y otros tipos
de pruebas tienen objetivos diferentes y se enfocan en distintas partes de la aplicación. Aquí está una descripción de
las diferencias clave entre ellos:

### Tests de Casos de Uso (Usecases)

**Objetivo**: Los tests de casos de uso verifican la lógica de negocio de tu aplicación. Se aseguran de que los flujos
de negocio definidos en los casos de uso funcionen correctamente.

**Enfoque**:

* Lógica de Negocio: Se centran en la lógica de negocio y los flujos de trabajo que ocurren cuando se ejecutan los casos
  de uso.
* Interacción de Componentes: A menudo implican la interacción de múltiples componentes (repositorios, servicios de
  notificación, etc.).
* Pruebas de Integración: Aunque los tests de casos de uso suelen ser considerados tests unitarios, también pueden ser
  considerados tests de integración ligeros ya que verifican la interacción entre componentes.
* Mocks y Stubs: Frecuentemente utilizan mocks o stubs para aislar la lógica de negocio y evitar dependencias en
  componentes externos como bases de datos reales.

### Tests de Repositorio (Repositories)

**Objetivo**: Los tests de repositorio verifican las operaciones de persistencia de datos. Se aseguran de que las
interacciones con la base de datos funcionen correctamente.

**Enfoque**:

* Operaciones CRUD: Se centran en las operaciones de creación, lectura, actualización y eliminación (CRUD) en la base de
  datos.
* Consultas de Base de Datos: Verifican que las consultas y las interacciones con la base de datos sean correctas.
* Persistencia de Datos: Se aseguran de que los datos se guarden, recuperen y borren correctamente en la base de datos.
* Tests de Integración: A menudo son considerados tests de integración porque verifican la integración con la base de
  datos.

### Tests de Handlers (Controladores)

**Objetivo**: Los tests de handlers verifican que los endpoints de tu API funcionen correctamente. Se aseguran de que
las rutas respondan con los códigos de estado y datos esperados.

**Enfoque**:

* Rutas y Endpoints: Se centran en las rutas de tu API y en las respuestas HTTP.
* Interacción con Casos de Uso: Verifican que los handlers llamen a los casos de uso correctos y manejen las respuestas
  adecuadamente.
* Respuestas HTTP: Aseguran que las respuestas HTTP (códigos de estado, headers, cuerpo) sean correctas.
* Mocks y Stubs: Utilizan mocks o stubs para aislar los handlers y evitar dependencias en la lógica de negocio o en la
  base de datos.

### Tests Unitarios

**Objetivo**: Verificar que unidades individuales de código (funciones, métodos, clases) funcionen correctamente.

**Enfoque**:

* Pequeñas unidades de código: Se centran en funciones y métodos individuales.
* Mocks y Stubs: Utilizan mocks y stubs para aislar la unidad de código.
* Pruebas de Integración: Ninguna, ya que están aislados.

### Tests de Integración

**Objetivo**: Verificar que múltiples componentes o sistemas funcionen correctamente juntos.

**Enfoque**:

* Interacción entre componentes: Verifican la interacción entre distintas partes del sistema (bases de datos, servicios
  externos, etc.).
* Mocks y Stubs: Pueden utilizarse, pero a menudo se prefiere probar con componentes reales.
* Pruebas de Integración: Completas, ya que prueban la interacción entre componentes.

### Tests Funcionales (End-to-End)

**Objetivo**: Verificar que la aplicación funcione correctamente desde el punto de vista del usuario final.

**Enfoque**:

* Flujos completos de la aplicación: Desde la entrada del usuario hasta la respuesta final.
* Mocks y Stubs: Generalmente no se utilizan, ya que se desea probar el sistema completo.
* Pruebas de Integración: Completas, ya que prueban el sistema en su totalidad.

### Tests de Carga y Rendimiento

**Objetivo**: Verificar cómo se comporta la aplicación bajo carga y evaluar su rendimiento.

**Enfoque**:

* Pruebas de estrés y de carga: Evaluar el rendimiento bajo diferentes condiciones de carga.
* Mocks y Stubs: No se utilizan, se prueba el sistema real.
* Pruebas de Integración: Completas, ya que prueban el sistema bajo condiciones de carga.

### Tests de Seguridad

**Objetivo**: Identificar vulnerabilidades de seguridad en la aplicación.

**Enfoque**:

* Pruebas de penetración y análisis de vulnerabilidades: Evaluar la seguridad del sistema.
* Mocks y Stubs: Generalmente no se utilizan, ya que se desea probar el sistema real.
* Pruebas de Integración: Completas, ya que prueban el sistema en su totalidad.

### Comparación

| Aspecto                        | Tests de Casos de Uso                 | Tests de Repositorio                            | Tests de Handlers                       | Tests Unitarios              | Tests de Integración       | Tests Funcionales (E2E)      | Tests de Carga y Rendimiento | Tests de Seguridad           |
|--------------------------------|---------------------------------------|-------------------------------------------------|-----------------------------------------|------------------------------|----------------------------|------------------------------|------------------------------|------------------------------|
| **Objetivo**                   | Verificar la lógica de negocio        | Verificar la persistencia de datos              | Verificar la respuesta de los endpoints | Verificar unidades de código | Verificar interacción      | Verificar flujos completos   | Evaluar rendimiento          | Identificar vulnerabilidades |
| **Enfoque**                    | Lógica de negocio y flujos de trabajo | Operaciones CRUD y consultas de base de datos   | Rutas y respuestas HTTP                 | Pequeñas unidades de código  | Interacción de componentes | Flujos completos del sistema | Pruebas de estrés y carga    | Pruebas de penetración       |
| **Interacción de Componentes** | Sí                                    | No (foco en interacciones con la base de datos) | Sí (llama a casos de uso)               | No                           | Sí                         | Sí                           | Sí                           | Sí                           |
| **Mocks y Stubs**              | Usualmente sí                         | Generalmente no                                 | Usualmente sí                           | Sí                           | Pueden utilizarse          | Generalmente no              | No                           | No                           |
| **Pruebas de Integración**     | Ligeras                               | Completas                                       | Ligeras a completas                     | Ninguna                      | Completas                  | Completas                    | Completas                    | Completas                    |

### ¿Cuándo Usar Cada Uno?

* **Tests de Casos de Uso**: Úsalos para asegurarte de que tu lógica de negocio funcione correctamente,
  independientemente de los detalles de implementación de la persistencia de datos.
    * **Ejemplo**: Crear un usuario, validar la lógica de negocio que asegura que el usuario tiene un correo electrónico
      único.
* **Tests de Repositorio**: Úsalos para asegurarte de que tus operaciones de base de datos sean correctas y para
  verificar que los datos se gestionen correctamente en la base de datos.
    * **Ejemplo**: Insertar un usuario en la base de datos, obtener un usuario por ID desde la base de datos.
* **Tests de Handlers**: Úsalos para asegurarte de que los endpoints de tu API respondan correctamente, verificando las
  rutas, los códigos de estado HTTP y los datos de respuesta.
    * **Ejemplo**: Verificar que un endpoint para crear un usuario responde con un código de estado 201 Created y que el
      cuerpo de la respuesta contiene los datos del usuario creado.
* **Tests Unitarios**: Úsalos para verificar que las unidades individuales de código funcionen correctamente.
    * **Ejemplo**: Probar funciones y métodos individuales en cualquier capa de la arquitectura.
* **Tests de Integración**: Úsalos para verificar que múltiples componentes o sistemas funcionen correctamente juntos.
    * **Ejemplo**: Probar la interacción entre la capa de aplicación y la capa de infraestructura.
* **Tests Funcionales (End-to-End)**: Úsalos para verificar el funcionamiento completo de la aplicación desde el punto
  de vista del usuario final.
    * **Ejemplo**: Simular el registro de un usuario y verificar que el flujo completo (desde la solicitud HTTP hasta la
      persistencia en la base de datos) funcione correctamente.
* **Tests de Carga y Rendimiento**: Úsalos para evaluar el comportamiento de la aplicación bajo condiciones de carga y
  su rendimiento.
    * **Ejemplo**: Verificar cómo responde la API bajo una alta carga de solicitudes.
* **Tests de Seguridad**: Úsalos para identificar vulnerabilidades de seguridad en la aplicación.
    * **Ejemplo**: Realizar pruebas de penetración para encontrar posibles fallos de seguridad en los endpoints de la
      API.

### Resumen

* **Tests de Casos de Uso**: Verifican que la lógica de negocio funcione correctamente.
* **Tests de Repositorio**: Aseguran que las operaciones de base de datos sean correctas.
* **Tests de Handlers**: Verifican que los endpoints de tu API respondan correctamente.
* **Tests Unitarios**: Verifican que las unidades individuales de código funcionen correctamente.
* **Tests de Integración**: Aseguran que los componentes funcionen correctamente juntos.
* **Tests Funcionales (End-to-End)**: Verifican el funcionamiento completo del sistema desde la perspectiva del usuario
  final.
* **Tests de Carga y Rendimiento**: Evaluan el comportamiento de la aplicación bajo condiciones de carga y rendimiento.
* **Tests de Seguridad**: Identifican vulnerabilidades de seguridad en la aplicación.

Todos estos tipos de tests son importantes y complementarios. Los tests de casos de uso aseguran que tu lógica de
negocio funcione correctamente, los tests de repositorio aseguran que los datos se persisten y se recuperan
correctamente, los tests de handlers aseguran que los endpoints de tu API respondan correctamente, los tests unitarios
aseguran que las unidades individuales de código funcionen correctamente, los tests de integración aseguran que los
componentes funcionen correctamente juntos, los tests funcionales aseguran el funcionamiento completo del sistema, los
tests de carga y rendimiento evalúan el comportamiento bajo condiciones de carga, y los tests de seguridad identifican
vulnerabilidades. Al tener todos estos tests, puedes estar seguro de que tu aplicación está bien probada y es robusta.
