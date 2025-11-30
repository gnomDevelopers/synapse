package entities

type Assignment struct {
	ID           int      `json:"id" db:"id"`
	Name         string   `json:"name" db:"name"`
	Deadline     string   `json:"deadline" db:"deadline"`
	Description  string   `json:"description" db:"description"`
	Files        []string `json:"files" db:"files"`
	DisciplineID int      `json:"discipline_id" db:"discipline_id"`
}

type CreateAssignment struct {
	Name         string   `json:"name" db:"name"`
	Deadline     string   `json:"deadline" db:"deadline"`
	Description  string   `json:"description" db:"description"`
	Files        []string `json:"files" db:"files"`
	DisciplineID int      `json:"discipline_id" db:"discipline_id"`
}

type UpdateAssignment struct {
	Name         string   `json:"name" db:"name"`
	Deadline     string   `json:"deadline" db:"deadline"`
	Description  string   `json:"description" db:"description"`
	Files        []string `json:"files" db:"files"`
	DisciplineID int      `json:"discipline_id" db:"discipline_id"`
}
