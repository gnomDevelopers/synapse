package entities

type Discipline struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	TeacherID int    `json:"teacher_id" db:"teacher_id"`
}

type CreateDiscipline struct {
	Name string `json:"name" db:"name"`
}

type UpdateDiscipline struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
