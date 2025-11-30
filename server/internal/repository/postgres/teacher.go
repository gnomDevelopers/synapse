package postgres

import "synapse/internal/entities"

func (db *DB) CreateTeacher(fullName, email, password string) error {
	query := `INSERT INTO teachers (full_name, email, password) VALUES ($1, $2, $3)`
	_, err := db.DB.Exec(query, fullName, email, password)
	return err
}

func (db *DB) GetTeacherByEmail(email string) (*entities.Teacher, error) {
	query := `SELECT id, full_name, email, password FROM teachers WHERE email = $1`
	row := db.DB.QueryRow(query, email)

	var teacher entities.Teacher
	err := row.Scan(&teacher.ID, &teacher.FullName, &teacher.Email, &teacher.Password)
	if err != nil {
		return nil, err
	}
	return &teacher, nil
}
