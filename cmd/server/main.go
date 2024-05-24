package main

import (
	"github.com/drossan/core-api/adapters"
	"github.com/drossan/core-api/config"
	_ "github.com/drossan/core-api/docs"
	"github.com/drossan/core-api/infrastructure/db"
	"github.com/drossan/core-api/infrastructure/router"
	"github.com/drossan/core-api/interfaces/api"
	"github.com/drossan/core-api/middleware"
	"github.com/drossan/core-api/seeder"
	"github.com/drossan/core-api/service"
	"github.com/drossan/core-api/usecase"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
)

// @title Core API
// @version 1.0
// @description Core for Intranet API in Go using Echo and Gorm.
// @termsOfService https://www.drossan.com/license

// @contact.name API Support
// @contact.url https://www.drossan.com
//@contact.email hola@drossan.com

// @license.name Proprietary
// @license.url https://www.drossan.com/license

// @host localhost:8080
// @BasePath /
func main() {
	// Cargar configuración
	cfg := config.LoadConfig()

	// Inicializar base de datos
	dbConn := db.Initialize(cfg.Database.URL)

	// Migrar base de datos
	db.Migrate(dbConn)

	// Crear el servicio de notificaciones
	notificationService := service.NewNotificationService()

	// Registrar notificador de email
	emailNotifier := adapters.NewEmailNotifier(
		cfg.Email.SMTPHost,
		cfg.Email.SMTPPort,
		cfg.Email.SMTPUser,
		cfg.Email.SMTPPassword,
		cfg.Email.FromEmail,
	)
	notificationService.RegisterNotifier("email", emailNotifier)

	// Registrar notificador de Slack
	slackNotifier := adapters.NewSlackNotifier()
	notificationService.RegisterNotifier("slack", slackNotifier)

	// Sembrar datos
	seeder.Seed()

	// Crear instancias de los repositorios
	userRepo := db.NewUserRepository(dbConn)
	formRepo := db.NewFormRepository(dbConn)
	levelRepo := db.NewLevelRepository(dbConn)
	levelPrivilegesRepo := db.NewLevelPrivilegesRepository(dbConn)
	menuTreeRepo := db.NewMenuTreeRepository(dbConn)

	// Inicializar casos de uso
	userUseCase := usecase.NewUserUseCase(userRepo)
	formUseCase := usecase.NewFormUseCase(formRepo)
	levelUseCase := usecase.NewLevelUseCase(levelRepo)
	levelPrivilegesUseCase := usecase.NewLevelPrivilegesUseCase(levelPrivilegesRepo)
	menuTreeUseCase := usecase.NewMenuTreeUseCase(menuTreeRepo)

	// Iniciar rutas
	e, r, a, prefix := router.NewEchoRouter(cfg.Server.JWTSecret)
	r.Use(middleware.NewAuthorizationMiddleware(levelRepo, formRepo, levelPrivilegesRepo, prefix))

	// Inicializar manejadores y registrar rutas
	userHandler := api.NewUserHandler(e, userUseCase)
	formHandler := api.NewFormHandler(e, formUseCase)
	levelHandler := api.NewLevelHandler(e, levelUseCase)
	levelPrivilegesHandler := api.NewLevelPrivilegesHandler(e, levelPrivilegesUseCase)
	menuTreeHandler := api.NewMenuTreeHandler(e, menuTreeUseCase)

	// Registro de rutas
	userHandler.AuthRoutes(a)
	userHandler.RegisterRoutes(r)
	formHandler.RegisterRoutes(r)
	levelHandler.RegisterRoutes(r)
	levelPrivilegesHandler.RegisterRoutes(r)
	menuTreeHandler.RegisterRoutes(r)

	if cfg.App.Env != "production" {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	// Enviar notificación a Slack cuando se levanta el API
	err := notificationService.SendNotification(
		[]string{"slack"},
		"The API has been started successfully!",
	)
	if err != nil {
		log.Printf("Failed to send Slack notification: %v", err)
	}

	// Iniciar el servidor
	log.Fatal(e.Start(cfg.Server.Address))
}
