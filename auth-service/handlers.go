package main

import "github.com/gofiber/fiber/v2"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// @Summary		Registra um novo usuário
// @Description	Cria um novo usuário com username e password
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			user	body		LoginRequest	true	"Dados do Usuário"
// @Success		201		{object}	User
// @Failure		400		{object}	map[string]string
// @Router			/register [post]
func register(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	hashedPassword, _ := HashPassword(user.Password)
	user.Password = hashedPassword
	DB.Create(&user)

	return c.JSON(user)
}

// @Summary		Faz login do usuário
// @Description	Autentica um usuário e retorna access e refresh tokens
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			login	body		LoginRequest	true	"Credenciais do Usuário"
// @Success		200		{object}	TokenResponse
// @Failure		401		{object}	map[string]string
// @Router			/login [post]
func login(c *fiber.Ctx) error {
	input := new(User)
	if err := c.BodyParser(input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var user User
	DB.Where("username = ?", input.Username).First(&user)
	if user.ID == 0 || !CheckPasswordHash(input.Password, user.Password) {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, _ := GenerateJWT(user.ID)
	return c.JSON(fiber.Map{"token": token})
}
