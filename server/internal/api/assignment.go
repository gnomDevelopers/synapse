package handler

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"path/filepath"
	"strconv"
	"synapse/internal/entities"
	"synapse/util"
	"time"
)

// CreateAssignment
// @Tags         Assignment
// @Summary      Создание нового задания с загрузкой файлов
// @Accept       multipart/form-data
// @Produce      json
// @Param        name         formData string true  "Название задания"
// @Param        description  formData string true  "Описание задания"
// @Param        deadline     formData string true  "Срок сдачи"
// @Param        discipline_id formData int   true  "ID дисциплины"
// @Param        media        formData []file   false "Файлы задания" collectionFormat(multi)
// @Success      200 {object} entities.Assignment "Созданное задание"
// @Failure      400 {object} entities.Error "Некорректный запрос или ошибка загрузки файлов"
// @Failure      401 {object} entities.Error "Ошибка аутентификации"
// @Failure      403 {object} entities.Error "Недостаточно прав"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/assignment [post]
// @Security ApiKeyAuth
func (h *Handler) CreateAssignment(c *fiber.Ctx) error {
	var e entities.CreateAssignment

	// Поля формы
	e.Name = c.FormValue("name")
	e.Description = c.FormValue("description")
	e.Deadline = c.FormValue("deadline")
	e.DisciplineID, _ = strconv.Atoi(c.FormValue("discipline_id"))

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{
			Error: "cannot parse multipart form: " + err.Error(),
		})
	}

	// ---- MULTI FILE UPLOAD ----
	media := form.File["media"]
	e.Files = make([]string, 0, len(media)) // массив имён файлов

	for _, file := range media {
		ext := filepath.Ext(file.Filename)
		randomFileName := util.GenerateRandomFileName(ext)
		tmpPath := fmt.Sprintf("./tmp/%s", randomFileName)

		// 1. Сохраняем временно
		if err := c.SaveFile(file, tmpPath); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(entities.Error{
				Error: fmt.Sprintf("error saving file %s: %s", file.Filename, err.Error()),
			})
		}

		// 2. Загружаем в MinIO с уникальным именем
		err = h.repository.S3.FPutObject(
			context.Background(),
			"assignment-media",
			randomFileName,
			tmpPath,
			file.Header.Get("Content-Type"),
		)
		if err != nil {
			os.Remove(tmpPath)
			return c.Status(fiber.StatusBadRequest).JSON(entities.Error{
				Error: fmt.Sprintf("error uploading %s to minio: %s", file.Filename, err.Error()),
			})
		}

		// 3. Удаляем временный файл
		if err := os.Remove(tmpPath); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(entities.Error{
				Error: fmt.Sprintf("error removing tmp file %s: %s", file.Filename, err.Error()),
			})
		}

		// 4. Добавляем уникальное имя файла в список
		e.Files = append(e.Files, randomFileName)
	}

	// ---- СОХРАНЯЕМ ASSIGNMENT В БД ----
	if err := h.repository.DB.CreateAssignment(e); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{
			Error: err.Error(),
		})
	}

	return c.Status(200).JSON(e)
}

// GetAssignmentByID
// @Tags         Assignment
// @Summary      Получение задания по ID с presigned ссылками на файлы
// @Accept       json
// @Produce      json
// @Param        id   path int true "ID задания"
// @Success      200 {object} entities.Assignment "Задание с файлами"
// @Failure      400 {object} entities.Error "Некорректный запрос"
// @Failure      401 {object} entities.Error "Ошибка аутентификации"
// @Failure      403 {object} entities.Error "Недостаточно прав"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/assignment/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) GetAssignmentByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{
			Error: "invalid assignment id",
		})
	}

	assignment, err := h.repository.DB.GetAssignmentByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{
			Error: err.Error(),
		})
	}

	// ---- Генерируем presigned URL для каждого файла ----
	filesWithURL := make([]string, 0, len(assignment.Files))
	for _, fileName := range assignment.Files {
		url, err := h.repository.S3.PresignedGetObject(
			c.Context(),
			"assignment-media",
			fileName,
			24*time.Hour, //presigned срок действия
		)
		if err != nil {
			fmt.Printf("error generating presigned URL for %s: %s\n", fileName, err.Error())
			continue
		}
		filesWithURL = append(filesWithURL, url.String())
	}
	assignment.Files = filesWithURL

	return c.Status(fiber.StatusOK).JSON(assignment)
}

// GetAssignmentsByDiscipline
// @Tags         Assignment
// @Summary      Получение всех заданий по ID дисциплины с presigned ссылками
// @Accept       json
// @Produce      json
// @Param        id   path int true "ID дисциплины"
// @Success      200 {array} entities.Assignment "Список заданий с файлами"
// @Failure      400 {object} entities.Error "Некорректный запрос"
// @Failure      401 {object} entities.Error "Ошибка аутентификации"
// @Failure      403 {object} entities.Error "Недостаточно прав"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/discipline/{id}/assignment [get]
// @Security ApiKeyAuth
func (h *Handler) GetAssignmentsByDiscipline(c *fiber.Ctx) error {
	disciplineID, err := c.ParamsInt("id")
	if err != nil || disciplineID <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{
			Error: "invalid discipline id",
		})
	}

	assignments, err := h.repository.DB.GetAssignmentsByDiscipline(disciplineID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{
			Error: err.Error(),
		})
	}

	// Генерируем presigned URL для файлов каждого задания
	for i := range assignments {
		filesWithURL := make([]string, 0, len(assignments[i].Files))
		for _, fileName := range assignments[i].Files {
			url, err := h.repository.S3.PresignedGetObject(
				c.Context(),
				"assignment-media", // используем один и тот же бакет
				fileName,
				time.Hour,
			)
			if err != nil {
				fmt.Printf("error generating presigned URL for %s: %s\n", fileName, err.Error())
				continue
			}
			filesWithURL = append(filesWithURL, url.String())
		}
		assignments[i].Files = filesWithURL
	}

	return c.Status(fiber.StatusOK).JSON(assignments)
}
