package repository

import (
	minios3 "synapse/internal/repository/minio"
	"synapse/internal/repository/postgres"
)

type Repository struct {
	DB *postgres.DB
	S3 *minios3.S3
}
