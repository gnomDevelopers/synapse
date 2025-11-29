package entities

type Question struct {
	ID               int      `json:"id" db:"id"`
	TestID           int      `json:"test_id" db:"test_id"`
	QuestionText     string   `json:"question_text" db:"question_text"`
	CorrectAnswers   []string `json:"correct_answers" db:"correct_answers"`
	IncorrectAnswers []string `json:"incorrect_answers" db:"incorrect_answers"`
}
