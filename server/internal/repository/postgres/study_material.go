package postgres

import (
	"synapse/internal/entities"
)

func (db *DB) CreateStudyMaterial(teacherID int, name, link, fileName string) error {
	query := `INSERT INTO study_materials (teacher_id, name, link, file_name) VALUES ($1, $2, $3, $4)`
	_, err := db.DB.Exec(query, teacherID, name, link, fileName)
	return err
}

func (db *DB) GetStudyMaterialsByTeacher(teacherID int) ([]entities.StudyMaterial, error) {
	query := `SELECT id, teacher_id, name, link, file_name FROM study_materials WHERE teacher_id = $1 ORDER BY id`
	rows, err := db.DB.Query(query, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var materials []entities.StudyMaterial
	for rows.Next() {
		var material entities.StudyMaterial
		err := rows.Scan(&material.ID, &material.TeacherID, &material.Name, &material.Link, &material.FileName)
		if err != nil {
			return nil, err
		}
		materials = append(materials, material)
	}
	return materials, nil
}

func (db *DB) GetStudyMaterialByID(id int) (*entities.StudyMaterial, error) {
	query := `SELECT id, teacher_id, name, link, file_name FROM study_materials WHERE id = $1`
	row := db.DB.QueryRow(query, id)

	var material entities.StudyMaterial
	err := row.Scan(&material.ID, &material.TeacherID, &material.Name, &material.Link, &material.FileName)
	if err != nil {
		return nil, err
	}
	return &material, nil
}

func (db *DB) UpdateStudyMaterial(id int, teacherID int, name, link, fileName string) error {
	query := `UPDATE study_materials SET name = $1, link = $2, file_name = $3 WHERE id = $4 AND teacher_id = $5`
	_, err := db.DB.Exec(query, name, link, fileName, id, teacherID)
	return err
}

func (db *DB) DeleteStudyMaterial(id int) error {
	query := `DELETE FROM study_materials WHERE id = $1`
	_, err := db.DB.Exec(query, id)
	return err
}
