package entities

type Test struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Deadline string `json:"deadline" db:"deadline"`
}
