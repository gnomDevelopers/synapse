package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"synapse/internal/entities"
	"synapse/util"
)

// SignUp
// @Tags         Auth
// @Summary      Регистрация пользователя
// @Accept       json
// @Produce      json
// @Param        request body entities.CreateTeacher true "Данные пользователя"
// @Success      200 {object} entities.Response "Пользователь успешно зарегистрирован"
// @Failure      400 {object} entities.Error "Некорректный запрос"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /sign-up [post]
func (h *Handler) SignUp(c *fiber.Ctx) error {
	var u entities.CreateTeacher
	if err := c.BodyParser(&u); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	u.Password = hashedPassword

	fmt.Println("call h.repository.DB.CreateTeacher")
	err = h.repository.DB.CreateTeacher(u.FullName, u.Email, u.Password)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Println("success")
	return c.Status(fiber.StatusOK).JSON(entities.Response{Response: "successful creating teacher"})
}

// Login
// @Tags         Auth
// @Summary      Аутентификация пользователя
// @Description  Вход пользователя в систему с выдачей токенов доступа и обновления
// @Accept       json
// @Produce      json
// @Param        request body entities.LoginUserRequest true "Данные для входа"
// @Success      200 {object} entities.LoginUserResponse "Успешная аутентификация"
// @Failure      400 {object} entities.Error "Некорректные данные для входа"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	var user entities.LoginUserRequest
	if err := c.BodyParser(&user); err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Println("call h.repository.DB.DBUserGetByEmail")
	u, err := h.repository.DB.GetTeacherByEmail(user.Email)
	if err != nil {
		fmt.Println("wrong data")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "wrong data"})
	}

	fmt.Println("call util.CheckPassword")
	err = util.CheckPassword(user.Password, u.Password)
	if err != nil {
		fmt.Println("wrong data")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "wrong data"})
	}

	TokenExpiration := 100
	fmt.Println("call pkg.GenerateAccessToken")
	accessToken, err := h.GenerateAccessToken(u.ID, TokenExpiration)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Println("call pkg.GenerateRefreshToken")
	refreshToken, err := h.GenerateRefreshToken(u.ID)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	res := entities.LoginUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	fmt.Println("success")
	return c.Status(fiber.StatusOK).JSON(res)
}
