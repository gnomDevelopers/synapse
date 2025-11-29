package entities

type Grade struct {
	ID         int    `json:"id" db:"id"`
	SolutionID int    `json:"solution_id" db:"solution_id"`
	GradeValue int    `json:"grade_value" db:"grade_value"`
	Comment    string `json:"comment" db:"comment"`
}
