package handler

import (
	_ "synapse/docs"

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

	f.Post("/sign-up", h.SignUp)
	f.Post("/login", h.Login)

	authGroup := f.Group("/auth")
	authGroup.Use(func(c *fiber.Ctx) error {
		return h.WithJWTAuth(c)
	})

	authGroup.Post("/discipline", h.CreateDiscipline)
	authGroup.Get("/discipline", h.GetDisciplinesByTeacher)
	authGroup.Put("/discipline", h.UpdateDiscipline)
	authGroup.Delete("/discipline/:id", h.DeleteDiscipline)

	authGroup.Post("/study-material", h.CreateStudyMaterial)
	authGroup.Get("/study-materials", h.GetStudyMaterialsByTeacher)
	authGroup.Get("/study-material/id/:id", h.GetStudyMaterialByID)
	authGroup.Put("/study-material", h.UpdateStudyMaterial)
	authGroup.Delete("/study-material/:id", h.DeleteStudyMaterial)

	return f
}
