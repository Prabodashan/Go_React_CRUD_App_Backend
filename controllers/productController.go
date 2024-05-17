package controllers

import (
	"server/database"
	"server/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Product{}, page))
}

func CreateProduct(c *fiber.Ctx) error {

	var Product models.Product

	if err := c.BodyParser(&Product); err != nil {
		return err
	}

	database.DB.Create(&Product)

	return c.JSON(Product)
}

func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	Product := models.Product{
		Id: uint(id),
	}

	database.DB.Find(&Product)

	return c.JSON(Product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	Product := models.Product{
		Id: uint(id),
	}

	if err := c.BodyParser(&Product); err != nil {
		return err
	}

	database.DB.Model(&Product).Updates(Product)

	return c.JSON(Product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	Product := models.Product{
		Id: uint(id),
	}

	database.DB.Delete(&Product)

	return nil
}
