package entities

type StudyMaterial struct {
	ID        int    `json:"id" db:"id"`
	TeacherID int    `json:"teacher_id" db:"teacher_id"`
	Name      string `json:"name" db:"name"`
	Link      string `json:"link" db:"link"`
	FileName  string `json:"file_name" db:"file_name"`
}

type CreateStudyMaterial struct {
	Name     string `json:"name" db:"name"`
	Link     string `json:"link" db:"link"`
	FileName string `json:"file_name" db:"file_name"`
}

type UpdateStudyMaterial struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Link     string `json:"link" db:"link"`
	FileName string `json:"file_name" db:"file_name"`
}
