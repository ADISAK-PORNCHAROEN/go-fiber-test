package handlers

import (
	"fiber-test/internal/models"
	"fiber-test/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// GetAllNames ดึงรายชื่อทั้งหมด
// @Summary      ดึงรายชื่อ (Get Names)
// @Description  ดึงข้อมูล Names ทั้งหมดในระบบ (ต้อง Login)
// @Tags         names
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.Name  "คืนค่าเป็น Array ของ Name"
// @Router       /api/v1/names [get]
func GetAllNames(c fiber.Ctx) error {
	q := c.Query("q")
	res, err := services.Search(q)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Get All Success",
		"data":    res,
	})
}

// GetNameById ดึงรายชื่อตาม ID
// @Summary      ดึงรายชื่อตาม ID (Get Name By Id)
// @Description  ดึงข้อมูล Name ตาม ID ที่ระบุ (ต้อง Login)
// @Tags         names
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Name ID"
// @Success      200  {object}  models.Name
// @Failure      404  {object}  map[string]interface{}
// @Router       /api/v1/names/{id} [get]
func GetNameById(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(400)
	}
	res, err := services.SearchNameById(id)
	if err != nil {
		return c.Status(404).JSON(err)
	}
	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    res,
	})
}

func CreateName(c fiber.Ctx) error {
	payload := new(models.Name)

	if err := c.Bind().Body(payload); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if payload.Fname == "" || payload.Lname == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "fname or lname is null",
		})
	}

	res, err := services.CreateName(*payload)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Create Success",
		"data":    res,
	})
}

func UpdateName(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	payload := new(models.Name)
	if err := c.Bind().Body(payload); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"messsage": err.Error(),
		})
	}

	res, err := services.UpdateName(id, *payload)
	if err != nil {
		return c.Status(404).JSON(err)
	}
	return c.JSON(fiber.Map{
		"message": "Update Success",
		"data":    res,
	})
}

func DeleteName(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res := services.DeleteName(id)
	if res == true {
		return c.JSON(fiber.Map{
			"message": "Delete Success",
		})
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func UploadFile(c fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// log.Println(file.Filename)

	err = c.SaveFile(file, "./uploads/"+file.Filename)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "File Upload Complete",
	})
}
