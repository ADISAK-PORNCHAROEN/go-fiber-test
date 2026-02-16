package namesService

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type Name struct {
	Id    int    `json:"id"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

var NameList []Name

func HelloWorld(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello World",
	})
}

func Names(c fiber.Ctx) error {
	res := c.Query("q")
	for _, name := range NameList {
		if res == name.Fname || res == name.Lname || res == strconv.Itoa(name.Id) {
			return c.JSON(fiber.Map{
				"message": "Search Success",
				"data":    name,
			})
		}
	}

	return c.JSON(fiber.Map{
		"message": "Get All Success",
		"data":    NameList,
	})
}

func NamesId(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(400)
	}

	for _, name := range NameList {
		if id == name.Id {
			return c.JSON(fiber.Map{
				"message": "Get Name By Id Success",
				"data":    name,
			})
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func CreateName(c fiber.Ctx) error {
	payload := new(Name)

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

	payload.Id = len(NameList) + 1
	NameList = append(NameList, *payload)

	return c.Status(201).JSON(fiber.Map{
		"message": "Create Success",
		"data":    payload,
	})
}

func UpdateName(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	payload := new(Name)
	if err := c.Bind().Body(payload); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"messsage": err.Error(),
		})
	}

	for i, name := range NameList {
		if name.Id == id {
			name.Fname = payload.Fname
			name.Lname = payload.Lname
			NameList[i] = name

			return c.JSON(fiber.Map{
				"message": i,
				"data":    name,
			})
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func DeleteName(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	for i, name := range NameList {
		if name.Id == id {
			NameList = append(NameList[:i], NameList[i+1:]...)

			return c.JSON(fiber.Map{
				"message": "Delete Success",
			})
		}
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

	log.Println(file.Filename)

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
