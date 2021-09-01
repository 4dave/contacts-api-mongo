package main

import (
	"contacts-api-mongo/entity"
	"contacts-api-mongo/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	return port
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/contact/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		contact, err := services.ContactGet(id)
		if err != nil {
			log.Fatal(err)
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(contact)
	})

	app.Get("/contacts", func(c *fiber.Ctx) error {
		contact, err := services.ContactGetAll()
		if err != nil {
			log.Fatal(err)
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(contact)
	})

	app.Post("/contact", func(c *fiber.Ctx) error {
		var contact entity.Contact
		err := c.BodyParser(&contact)
		if err != nil {
			c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		contact, _ = services.ContactCreate(contact)
		return c.Status(fiber.StatusOK).JSON(contact)
	})

	app.Delete("/contact/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := c.BodyParser(&id)
		if err != nil {
			c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		err = services.ContactDelete(id)
		if err != nil {
			log.Fatal(err)
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Contact deleted",
		})
	})

	// update contact by id
	app.Put("/contact/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var contact entity.Contact
		err := c.BodyParser(&contact)
		if err != nil {
			c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		contact, _ = services.ContactUpdate(id, contact)
		return c.Status(fiber.StatusOK).JSON(contact)
	})

	port := getPort()
	app.Listen(":" + port)
}
