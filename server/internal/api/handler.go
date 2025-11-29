package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"synapse/internal/config"
	"synapse/internal/repository"
)

type Handler struct {
	repository *repository.Repository
	conf       *config.Config
}

func NewHandler(repository *repository.Repository, conf *config.Config) *Handler {
	return &Handler{repository: repository, conf: conf}
}

func (h *Handler) Router() *fiber.App {
	f := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
	})

	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		//AllowCredentials: true,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))
	//f.Use(log.RequestLogger(h.logger))

	f.Get("/swagger/*", fiberSwagger.WrapHandler)
	f.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("healthy")
	})

	authGroup := f.Group("/auth")
	authGroup.Use(func(c *fiber.Ctx) error {
		return h.WithJWTAuth(c)
	})

	authGroup.Post("/discipline", h.CreateDiscipline)
	authGroup.Get("/discipline", h.GetDisciplinesByTeacher)
	authGroup.Put("/discipline", h.UpdateDiscipline)
	authGroup.Delete("/discipline", h.DeleteDiscipline)

	return f
}
