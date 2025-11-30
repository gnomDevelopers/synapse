package postgres

import (
	"github.com/lib/pq"
	"synapse/internal/entities"
)

func (db *DB) CreateAssignment(a entities.CreateAssignment) error {
	query := `
        INSERT INTO assignments (name, deadline, description, files, discipline_id)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err := db.DB.Exec(query,
		a.Name,
		a.Deadline,
		a.Description,
		pq.Array(a.Files),
		a.DisciplineID,
	)
	return err
}

func (db *DB) GetAssignmentsByDiscipline(disciplineID int) ([]entities.Assignment, error) {
	query := `
        SELECT id, name, deadline, description, files, discipline_id
        FROM assignments
        WHERE discipline_id = $1
        ORDER BY id
    `

	rows, err := db.DB.Query(query, disciplineID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var arr []entities.Assignment
	for rows.Next() {
		var a entities.Assignment
		err := rows.Scan(
			&a.ID,
			&a.Name,
			&a.Deadline,
			&a.Description,
			pq.Array(&a.Files),
			&a.DisciplineID,
		)
		if err != nil {
			return nil, err
		}
		arr = append(arr, a)
	}
	return arr, nil
}

func (db *DB) GetAssignmentByID(id int) (*entities.Assignment, error) {
	query := `
        SELECT id, name, deadline, description, files, discipline_id
        FROM assignments
        WHERE id = $1
        LIMIT 1
    `

	var a entities.Assignment
	err := db.DB.QueryRow(query, id).Scan(
		&a.ID,
		&a.Name,
		&a.Deadline,
		&a.Description,
		pq.Array(&a.Files),
		&a.DisciplineID,
	)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (db *DB) UpdateAssignment(id int, a entities.UpdateAssignment) error {
	query := `
        UPDATE assignments
        SET name = $1, deadline = $2, description = $3, files = $4, discipline_id = $5
        WHERE id = $6
    `

	_, err := db.DB.Exec(query,
		a.Name,
		a.Deadline,
		a.Description,
		pq.Array(a.Files),
		a.DisciplineID,
		id,
	)

	return err
}

func (db *DB) DeleteAssignment(id int) error {
	query := `DELETE FROM assignments WHERE id = $1`
	_, err := db.DB.Exec(query, id)
	return err
}
