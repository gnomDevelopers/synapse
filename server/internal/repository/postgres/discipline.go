package postgres

import (
	"database/sql"
	"synapse/internal/entities"
)

func CreateDiscipline(db *sql.DB, name string, teacherID int) error {
	query := `INSERT INTO disciplines (name, teacher_id) VALUES ($1, $2)`
	_, err := db.Exec(query, name, teacherID)
	return err
}

func GetDiscipline(db *sql.DB, id int) (*entities.Discipline, error) {
	query := `SELECT id, name, teacher_id FROM disciplines WHERE id = $1`
	row := db.QueryRow(query, id)

	var discipline entities.Discipline
	err := row.Scan(&discipline.ID, &discipline.Name, &discipline.TeacherID)
	if err != nil {
		return nil, err
	}
	return &discipline, nil
}

func GetDisciplinesByTeacher(db *sql.DB, teacherID int) ([]entities.Discipline, error) {
	query := `SELECT id, name, teacher_id FROM disciplines WHERE teacher_id = $1 ORDER BY id`
	rows, err := db.Query(query, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var disciplines []entities.Discipline
	for rows.Next() {
		var discipline entities.Discipline
		err := rows.Scan(&discipline.ID, &discipline.Name, &discipline.TeacherID)
		if err != nil {
			return nil, err
		}
		disciplines = append(disciplines, discipline)
	}
	return disciplines, nil
}

func GetAllDisciplines(db *sql.DB) ([]entities.Discipline, error) {
	query := `SELECT id, name, teacher_id FROM disciplines ORDER BY id`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var disciplines []entities.Discipline
	for rows.Next() {
		var discipline entities.Discipline
		err := rows.Scan(&discipline.ID, &discipline.Name, &discipline.TeacherID)
		if err != nil {
			return nil, err
		}
		disciplines = append(disciplines, discipline)
	}
	return disciplines, nil
}

func UpdateDiscipline(db *sql.DB, id int, name string, teacherID int) error {
	query := `UPDATE disciplines SET name = $1, teacher_id = $2 WHERE id = $3`
	_, err := db.Exec(query, name, teacherID, id)
	return err
}

func DeleteDiscipline(db *sql.DB, id int) error {
	query := `DELETE FROM disciplines WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}
