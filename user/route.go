package user

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitUserRoute(router fiber.Router, db *gorm.DB) {
	db.AutoMigrate(&User{})
	// this will apply "/user" prefix to all route
	userRoute := router.Group("/user")
	userRoute.Get("/:id", func(c *fiber.Ctx) error { return getUserById(c, db) })
	userRoute.Post("/", func(c *fiber.Ctx) error { return createUser(c, db) })
	userRoute.Delete("/:id", func(c *fiber.Ctx) error { return delete(c, db) })
}
