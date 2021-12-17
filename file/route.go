package file

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/nu7hatch/gouuid"
	"gorm.io/gorm"
)

func InitImageRoute(router fiber.Router, db *gorm.DB) {
	api := router.Group("/file")
	api.Post("/", func(c *fiber.Ctx) error {
		if form, err := c.MultipartForm(); err == nil {
			// => *multipart.Form

			// Get all files from "documents" key:
			files := form.File["files"]
			// => []*multipart.FileHeader

			// Loop through files:
			for _, file := range files {
				fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
				// => "tutorial.pdf" 360641 "application/pdf"

				// Save the files to disk:
				u, uErr := uuid.NewV4()
				if uErr != nil {
					return uErr
				}
				fileLocation := "./upload/" + u.String() + "_" + file.Filename
				if err := c.SaveFile(file, fileLocation); err != nil {
					return err
				}
				c.JSON(fiber.Map{
					"url": fileLocation,
				})
			}
			return err
		}
		return c.SendString("failed")
	})
}
