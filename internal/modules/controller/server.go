package controller

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jsoniter "github.com/json-iterator/go"
	"github.com/martirosharutyunyan/fraudnetic-test-task/docs"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/common"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/config"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/modules/repository"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/modules/service"
	"github.com/scylladb/gocqlx/v3"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

type ServerConfig struct {
	DB *gocqlx.Session
}

// NewServer godoc
// @SecurityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
func NewServer(conf *ServerConfig) *fiber.App {
	eventRepo := repository.NewEvent(conf.DB, common.NewEventTable())
	eventService := service.NewEvent(eventRepo)

	// app
	router := fiber.New(fiber.Config{
		JSONEncoder: jsoniter.Marshal,
		JSONDecoder: jsoniter.Unmarshal,
	})
	prometheus := fiberprometheus.New("my-service-name")
	prometheus.RegisterAt(router, "/metrics")
	router.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return config.GetEnv() == "development"
		},
	}))
	router.Use(logger.New())
	router.Use(recover.New())
	router.Use(prometheus.Middleware)

	// swagger config
	docs.SwaggerInfo.BasePath = "/api"
	router.Get(
		"/swagger/*",
		fiberSwagger.FiberWrapHandler(fiberSwagger.PersistAuthorization(true)),
	)

	// router group
	apiRouter := router.Group("/api")

	// controllers
	NewEvent(eventService).Register(apiRouter)

	return router
}
