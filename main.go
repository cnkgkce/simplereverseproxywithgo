package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// app.Get("/:name",func(c *fiber.Ctx) error {
	// 	return c.SendString(fmt.Sprintf("Hello, %s!",c.Params("name")))
	// })

	app.Get("/:key/*",ProxyHandler)

	app.Listen(":8000")


}