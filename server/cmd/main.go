package main

import (
	"github.com/joho/godotenv"
	zerolog "github.com/rs/zerolog/log"
	"log"
	handler "synapse/internal/api"
	"synapse/internal/config"
	"synapse/internal/logger"
	"synapse/internal/repository"
	minios3 "synapse/internal/repository/minio"
	"synapse/internal/repository/postgres"
	"synapse/util"
)

// @title synapse API
// @version 1.0
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}
	conf := config.NewEnvConfig()
	logger := logger.Setup(conf)
	zerolog.Logger = *logger

	util.CreateTmpDirectory()

	db, err := postgres.NewDatabase(conf)
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}
	s3 := minios3.MinioConnection(conf)

	repo := repository.Repository{DB: db, S3: s3}
	handlers := handler.NewHandler(&repo, conf)
	repo.CreateMocks()
	app := handlers.Router()
	app.Listen(":8080")
}
