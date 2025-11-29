package postgres

import (
	"synapse/internal/entities"
)

func (db *DB) CreateDiscipline(name string, teacherID int) error {
	query := `INSERT INTO disciplines (name, teacher_id) VALUES ($1, $2)`
	_, err := db.DB.Exec(query, name, teacherID)
	return err
}

func (db *DB) GetDisciplinesByTeacher(teacherID int) ([]entities.Discipline, error) {
	query := `SELECT id, name, teacher_id FROM disciplines WHERE teacher_id = $1 ORDER BY id`
	rows, err := db.DB.Query(query, teacherID)
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

func (db *DB) UpdateDiscipline(id int, name string, teacherID int) error {
	query := `UPDATE disciplines SET name = $1, teacher_id = $2 WHERE id = $3`
	_, err := db.DB.Exec(query, name, teacherID, id)
	return err
}

func (db *DB) DeleteDiscipline(id int) error {
	query := `DELETE FROM disciplines WHERE id = $1`
	_, err := db.DB.Exec(query, id)
	return err
}
