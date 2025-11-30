package repository

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (r *Repository) CreateTeachers() {
	var count int
	err := r.DB.DB.QueryRow("SELECT COUNT(*) FROM teachers").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		for _, teacher := range MockTeachers {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(teacher.Password), bcrypt.DefaultCost)
			if err != nil {
				log.Fatalln(err.Error())
			}

			query := `
				INSERT INTO teachers (full_name, email, password) 
				VALUES ($1, $2, $3)
			`

			_, err = r.DB.DB.Exec(query, teacher.FullName, teacher.Email, string(hashedPassword))
			if err != nil {
				log.Fatal(err)
			}
		}

		log.Println("teachers created successfully")
	} else {
		log.Println("Teachers already exist, skipping creation")
	}
}

func (r *Repository) CreateDisciplines() {
	var count int
	err := r.DB.DB.QueryRow("SELECT COUNT(*) FROM disciplines").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		for _, discipline := range MockDisciplines {
			query := `
				INSERT INTO disciplines (name, teacher_id) 
				VALUES ($1, $2)
			`

			_, err = r.DB.DB.Exec(query, discipline.Name, discipline.TeacherID)
			if err != nil {
				log.Fatal(err)
			}
		}

		log.Println("disciplines created successfully")
	} else {
		log.Println("Disciplines already exist, skipping creation")
	}
}

func (r *Repository) CreateStudyMaterials() {
	var count int
	err := r.DB.DB.QueryRow("SELECT COUNT(*) FROM study_materials").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		for _, material := range MockStudyMaterials {
			query := `
				INSERT INTO study_materials (teacher_id, name, link, file_name, date) 
				VALUES ($1, $2, $3, $4, $5)
			`

			_, err = r.DB.DB.Exec(query, material.TeacherID, material.Name, material.Link, material.FileName, material.Date)
			if err != nil {
				log.Fatal(err)
			}
		}

		log.Println("study materials created successfully")
	} else {
		log.Println("Study materials already exist, skipping creation")
	}
}

func (r *Repository) CreateMocks() {
	r.CreateTeachers()
	r.CreateDisciplines()
	r.CreateStudyMaterials()
}
