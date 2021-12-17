package user

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func getUserById(c *fiber.Ctx, db *gorm.DB) error {
	user := new(User)
	id := c.Params("id")
	db.Where("id = ?", id).First(&user)
	return c.JSON(user)
}

func createUser(c *fiber.Ctx, db *gorm.DB) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	db.Create(&user)
	return c.JSON(user)
}
func delete(c *fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id")
	result := db.Delete(&User{}, id)
	return c.JSON(result)
}
