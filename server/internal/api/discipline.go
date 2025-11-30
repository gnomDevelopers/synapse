package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"synapse/internal/entities"
)

// CreateDiscipline
// @Tags         Discipline
// @Summary      Получение заявок представителей на регистрацию
// @Accept       json
// @Produce      json
// @Param        request body entities.CreateDiscipline true "adwawd"
// @Success      200 {array} entities.Response ""
// @Failure      400 {object} entities.Error "Некорректный запрос"
// @Failure      401 {object} entities.Error "Ошибка аутентификации"
// @Failure      403 {object} entities.Error "Недостаточно прав"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/discipline [post]
// @Security ApiKeyAuth
func (h *Handler) CreateDiscipline(c *fiber.Ctx) error {
	userID, ok := c.Locals("id").(int)
	if !ok {
		fmt.Println("error: cant get id from token")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cant get id from token"})
	}

	var e entities.CreateDiscipline
	if err := c.BodyParser(&e); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.repository.DB.CreateDiscipline(e.Name, userID)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Println("success")
	return c.Status(fiber.StatusOK).JSON(entities.Response{Response: "successful creating discipline"})
}

// GetDisciplinesByTeacher
// @Tags         Discipline
// @Summary      Получение дисциплин преподавателя
// @Accept       json
// @Produce      json
// @Success      200 {object} []entities.Discipline ""
// @Failure      400 {object} entities.Error "Некорректный запрос"
// @Failure      401 {object} entities.Error "Ошибка аутентификации"
// @Failure      403 {object} entities.Error "Недостаточно прав"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/discipline [get]
// @Security ApiKeyAuth
func (h *Handler) GetDisciplinesByTeacher(c *fiber.Ctx) error {
	userID, ok := c.Locals("id").(int)
	if !ok {
		fmt.Println("error: cant get id from token")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cant get id from token"})
	}

	disciplines, err := h.repository.DB.GetDisciplinesByTeacher(userID)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Println("success")
	return c.Status(fiber.StatusOK).JSON(disciplines)
}

// UpdateDiscipline
// @Tags         Discipline
// @Summary      Обновление дисциплины
// @Accept       json
// @Produce      json
// @Param        body body entities.UpdateDiscipline true "Данные для обновления"
// @Success      200 {object} entities.Response ""
// @Failure      400 {object} entities.Error "Некорректный запрос"
// @Failure      401 {object} entities.Error "Ошибка аутентификации"
// @Failure      403 {object} entities.Error "Недостаточно прав"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/discipline [put]
// @Security ApiKeyAuth
func (h *Handler) UpdateDiscipline(c *fiber.Ctx) error {
	userID, ok := c.Locals("id").(int)
	if !ok {
		fmt.Println("error: cant get id from token")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cant get id from token"})
	}

	var e entities.UpdateDiscipline
	if err := c.BodyParser(&e); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.repository.DB.UpdateDiscipline(e.ID, e.Name, userID)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Println("success")
	return c.Status(fiber.StatusOK).JSON(entities.Response{Response: "successful updating discipline"})
}

// DeleteDiscipline
// @Tags         Discipline
// @Summary      Удаление дисциплины
// @Accept       json
// @Produce      json
// @Param id path int true "ID" example(1)
// @Success      200 {object} entities.Response ""
// @Failure      400 {object} entities.Error "Некорректный запрос"
// @Failure      401 {object} entities.Error "Ошибка аутентификации"
// @Failure      403 {object} entities.Error "Недостаточно прав"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/discipline/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) DeleteDiscipline(c *fiber.Ctx) error {
	disciplineID, err := c.ParamsInt("id")
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid discipline id"})
	}

	err = h.repository.DB.DeleteDiscipline(disciplineID)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Println("success")
	return c.Status(fiber.StatusOK).JSON(entities.Response{Response: "successful deleting discipline"})
}
