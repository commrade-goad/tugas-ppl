package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func (c *fiber.Ctx) error {
		return c.Redirect("/home")
        // return c.SendString("Hello, World!")
    })

	app.Get("/home", func (c *fiber.Ctx) error {
		return c.SendFile("./template/home.html", true)
	})

	app.Get("/about", func (c *fiber.Ctx) error {
		return c.SendFile("./template/about.html", true)
	})

    log.Fatal(app.Listen(":3000"))
}
