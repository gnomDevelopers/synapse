package postgres

const (
	createTableTeachers = `
CREATE TABLE IF NOT EXISTS teachers (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL
);`

	createTableDisciplines = `
CREATE TABLE IF NOT EXISTS disciplines (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    teacher_id INT REFERENCES teachers(id) ON DELETE CASCADE
);`

	createTableAssignments = `
CREATE TABLE IF NOT EXISTS assignments (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    deadline VARCHAR NOT NULL,
    description TEXT NOT NULL,
    files TEXT[],
    discipline_id INT REFERENCES disciplines(id) ON DELETE CASCADE
);`

	createTableSolutions = `
CREATE TABLE IF NOT EXISTS solutions (
id SERIAL PRIMARY KEY,
assignment_id INT REFERENCES assignments(id) ON DELETE CASCADE,
solution_text TEXT NOT NULL,
files TEXT[]
);`

	createTableGrades = `
CREATE TABLE IF NOT EXISTS grades (
id SERIAL PRIMARY KEY,
solution_id INT REFERENCES solutions(id) ON DELETE CASCADE,
grade_value INT NOT NULL,
comment TEXT
);`

	createTableTests = `
CREATE TABLE IF NOT EXISTS tests (
id SERIAL PRIMARY KEY,
name VARCHAR NOT NULL,
deadline VARCHAR NOT NULL
);`

	createTableQuestions = `
CREATE TABLE IF NOT EXISTS questions (
id SERIAL PRIMARY KEY,
test_id INT REFERENCES tests(id) ON DELETE CASCADE,
question_text TEXT NOT NULL,
correct_answers TEXT[],
incorrect_answers TEXT[]
);`
)
