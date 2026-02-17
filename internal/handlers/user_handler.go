package handlers

import (
	"fiber-test/internal/models"
	"fiber-test/internal/services"

	"github.com/gofiber/fiber/v3"
)

// Login เข้าสู่ระบบเพื่อรับ Token
// @Summary      Login
// @Description  ส่ง Username/Password เพื่อรับ JWT Token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body models.User true "ข้อมูล Login"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Router       /login [post]
func UserLogin(c fiber.Ctx) error {
	body := new(models.User)

	if err := c.Bind().Body(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if body.Username == "" || body.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Username or Password have space!",
		})
	}

	token, err := services.LoginService(body.Username, body.Password)

	if err != nil {
		return c.SendStatus(fiber.ErrUnauthorized.Code)
	}

	return c.JSON(fiber.Map{
		"message": "Login Success",
		"token":   token,
	})
}
