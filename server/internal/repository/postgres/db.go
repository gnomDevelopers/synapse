package postgres

import (
	"fmt"
	"log"
	"synapse/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	DB *sqlx.DB
}

func NewDatabase(envConf *config.Config) (*DB, error) {
	connectionString := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		envConf.Db.User,
		envConf.Db.Password,
		envConf.Db.Host,
		envConf.Db.Port,
		envConf.Db.Database,
	)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database connection: %v", err)
	}

	CreateTable(db)

	return &DB{db}, nil
}

func CreateTable(db *sqlx.DB) {
	//db.MustExec(createTableUsers)
	//db.MustExec(createTableUniversities)
	//db.MustExec(createTablePrograms)
	//db.MustExec(createTableUsersUniversities)
	//db.MustExec(createTableUsersPrograms)
	//db.MustExec(createTableDisciplines)
	//db.MustExec(createTableApplications)
	//db.MustExec(createTableApplicationDocuments)
	//db.MustExec(createTableReviews)
	//db.MustExec(createTableQuestions)
	//db.MustExec(createTableAnswers)
	//db.MustExec(createTableQuota)
	//db.MustExec(createTableTelegramMessageTemplates)
	//db.MustExec(createTableTestTasks)
	//db.MustExec(createTableTestTasksSolutions)
	//db.MustExec(createTableForms)
}
