package entities

type Solution struct {
	ID           int      `json:"id" db:"id"`
	AssignmentID int      `json:"assignment_id" db:"assignment_id"`
	SolutionText string   `json:"solution_text" db:"solution_text"`
	Files        []string `json:"files" db:"files"`
}
