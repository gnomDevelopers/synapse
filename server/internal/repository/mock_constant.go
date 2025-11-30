package repository

import "synapse/internal/entities"

var (
	MockTeachers = []entities.Teacher{
		{
			FullName: "Иванов Петр Сергеевич",
			Email:    "i.p.sergeevich@yandex.ru",
			Password: "password123",
		},
		{
			FullName: "Смирнова Анна Владимировна",
			Email:    "a.v.smirnova@yandex.ru",
			Password: "password123",
		},
		{
			FullName: "Козлов Дмитрий Иванович",
			Email:    "d.i.kozlov@yandex.ru",
			Password: "password123",
		},
	}

	MockDisciplines = []entities.Discipline{
		{
			Name:      "Математический анализ",
			TeacherID: 1,
		},
		{
			Name:      "Программирование на Go",
			TeacherID: 2,
		},
		{
			Name:      "Базы данных",
			TeacherID: 2,
		},
		{
			Name:      "Алгоритмы и структуры данных",
			TeacherID: 1,
		},
		{
			Name:      "Веб-разработка",
			TeacherID: 3,
		},
	}

	MockStudyMaterials = []entities.StudyMaterial{
		{
			TeacherID: 1,
			Name:      "Конспект лекций по матанализу",
			Link:      "/materials/matan_lectures.pdf",
			FileName:  "",
			Date:      "15.09.2024",
		},
		{
			TeacherID: 2,
			Name:      "Презентация по Go",
			Link:      "/materials/go_presentation.pptx",
			FileName:  "",
			Date:      "20.09.2024",
		},
		{
			TeacherID: 2,
			Name:      "Лабораторная работа по SQL",
			Link:      "",
			FileName:  "bd.docx",
			Date:      "10.09.2024",
		},
		{
			TeacherID: 1,
			Name:      "Задачи по алгоритмам",
			Link:      "/materials/algorithms_tasks.pdf",
			FileName:  "",
			Date:      "25.09.2024",
		},
		{
			TeacherID: 3,
			Name:      "Примеры веб-приложений",
			Link:      "/materials/web_examples.zip",
			FileName:  "",
			Date:      "05.09.2024",
		},
	}
)
