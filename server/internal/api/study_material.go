package handler

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"synapse/internal/entities"
	"time"
)

// CreateStudyMaterial
// @Tags         StudyMaterial
// @Summary      Создание учебного материала
// @Accept       json
// @Produce      json
// @Param name formData string true "Название"
// @Param link formData string true "Ссылка"
// @Param media formData file false "Загрузка файла"
// @Success      200 {object} entities.Response ""
// @Failure      400 {object} entities.Error "Некорректный запрос"
// @Failure      401 {object} entities.Error "Ошибка аутентификации"
// @Failure      403 {object} entities.Error "Недостаточно прав"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/study-material [post]
// @Security ApiKeyAuth
func (h *Handler) CreateStudyMaterial(c *fiber.Ctx) error {
	userID, ok := c.Locals("id").(int)
	if !ok {
		fmt.Println("error: cant get id from token")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cant get id from token"})
	}

	var e entities.CreateStudyMaterial
	if err := c.BodyParser(&e); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: err.Error()})
	}

	e.Date = time.Now().Format("02.01.2006")

	media := form.File["media"]
	if len(media) != 0 {
		e.FileName = media[0].Filename

		filePath := fmt.Sprintf("./tmp/%s", media[0].Filename)
		c.SaveFile(media[0], filePath)
		err = h.repository.S3.FPutObject(context.Background(), "study-material", media[0].Filename, filePath, media[0].Header.Get("Content-Type"))
		if err != nil {
			fmt.Printf("error creating file %s in minio: %s", media[0].Filename, err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error creating file %s in minio: %s", media[0].Filename, err.Error())})
		}
		err = os.Remove(filePath)
		if err != nil {
			fmt.Printf("error removing file %s: %s", media[0].Filename, err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error removing file %s: %s", media[0].Filename, err.Error())})
		}
	}

	err = h.repository.DB.CreateStudyMaterial(userID, e.Name, e.Link, e.FileName)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Println("success")
	return c.Status(fiber.StatusOK).JSON(entities.Response{Response: "successful creating study material"})
}

// GetStudyMaterialsByTeacher
// @Tags         StudyMaterial
// @Summary      Получение учебных материалов преподавателя
// @Accept       json
// @Produce      json
// @Success      200 {object} []entities.StudyMaterial ""
// @Failure      400 {object} entities.Error "Некорректный запрос"
// @Failure      401 {object} entities.Error "Ошибка аутентификации"
// @Failure      403 {object} entities.Error "Недостаточно прав"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/study-materials [get]
// @Security ApiKeyAuth
func (h *Handler) GetStudyMaterialsByTeacher(c *fiber.Ctx) error {
	userID, ok := c.Locals("id").(int)
	if !ok {
		fmt.Println("error: cant get id from token")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cant get id from token"})
	}

	materials, err := h.repository.DB.GetStudyMaterialsByTeacher(userID)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Println("success")
	return c.Status(fiber.StatusOK).JSON(materials)
}

// GetStudyMaterialByID
// @Tags         StudyMaterial
// @Summary      Получение учебного материала по ID
// @Accept       json
// @Produce      json
// @Param        id query int true "ID учебного материала"
// @Success      200 {object} entities.StudyMaterial ""
// @Failure      400 {object} entities.Error "Некорректный запрос"
// @Failure      401 {object} entities.Error "Ошибка аутентификации"
// @Failure      403 {object} entities.Error "Недостаточно прав"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/study-material/id/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) GetStudyMaterialByID(c *fiber.Ctx) error {
	materialID, err := c.ParamsInt("id")
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid study material id"})
	}

	material, err := h.repository.DB.GetStudyMaterialByID(materialID)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	url, _ := h.repository.S3.PresignedGetObject(context.Background(), "study-material", material.FileName, 7*24*time.Hour)
	material.FileName = url.String()

	fmt.Println("success")
	return c.Status(fiber.StatusOK).JSON(material)
}

// UpdateStudyMaterial
// @Tags         StudyMaterial
// @Summary      Обновление учебного материала
// @Accept       json
// @Produce      json
// @Param        body body entities.UpdateStudyMaterial true "Данные для обновления"
// @Success      200 {object} entities.Response ""
// @Failure      400 {object} entities.Error "Некорректный запрос"
// @Failure      401 {object} entities.Error "Ошибка аутентификации"
// @Failure      403 {object} entities.Error "Недостаточно прав"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/study-material [put]
// @Security ApiKeyAuth
func (h *Handler) UpdateStudyMaterial(c *fiber.Ctx) error {
	userID, ok := c.Locals("id").(int)
	if !ok {
		fmt.Println("error: cant get id from token")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cant get id from token"})
	}

	var e entities.UpdateStudyMaterial
	if err := c.BodyParser(&e); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.repository.DB.UpdateStudyMaterial(e.ID, userID, e.Name, e.Link, e.FileName)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Println("success")
	return c.Status(fiber.StatusOK).JSON(entities.Response{Response: "successful updating study material"})
}

// DeleteStudyMaterial
// @Tags         StudyMaterial
// @Summary      Удаление учебного материала
// @Accept       json
// @Produce      json
// @Param        id path int true "ID учебного материала"
// @Success      200 {object} entities.Response ""
// @Failure      400 {object} entities.Error "Некорректный запрос"
// @Failure      401 {object} entities.Error "Ошибка аутентификации"
// @Failure      403 {object} entities.Error "Недостаточно прав"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/study-material/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) DeleteStudyMaterial(c *fiber.Ctx) error {
	materialID, err := c.ParamsInt("id")
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid study material id"})
	}

	err = h.repository.DB.DeleteStudyMaterial(materialID)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Println("success")
	return c.Status(fiber.StatusOK).JSON(entities.Response{Response: "successful deleting study material"})
}
