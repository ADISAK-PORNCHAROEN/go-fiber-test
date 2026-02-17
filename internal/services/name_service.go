package services

import (
	"fiber-test/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

var NameList = []models.Name{
	{Id: 1, Fname: "John", Lname: "Doe"},
	{Id: 2, Fname: "Admin", Lname: "System"},
}

func Search(q string) ([]models.Name, error) {
	if q == "" {
		return NameList, nil
	}

	result := []models.Name{}

	for _, name := range NameList {
		if q == name.Fname || q == name.Lname || q == strconv.Itoa(name.Id) {
			result = append(result, name)
		}
	}
	return result, nil
}

func SearchNameById(id int) (models.Name, error) {
	for _, name := range NameList {
		if id == name.Id {
			return name, nil
		}
	}

	return models.Name{}, fiber.ErrNotFound
}

func CreateName(body models.Name) (models.Name, error) {
	body.Id = len(NameList) + 1
	NameList = append(NameList, body)

	return body, nil
}

func UpdateName(id int, body models.Name) (models.Name, error) {
	if body.Fname == "" || body.Lname == "" {
		return models.Name{}, fiber.ErrBadRequest
	}
	for i, name := range NameList {
		if name.Id == id {
			name.Fname = body.Fname
			name.Lname = body.Lname
			NameList[i] = name

			return name, nil
		}
	}
	return models.Name{}, fiber.ErrNotFound
}

func DeleteName(id int) bool {
	for i, name := range NameList {
		if name.Id == id {
			NameList = append(NameList[:1], NameList[i+1:]...)
			return true
		}
	}
	return false
}
