package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type (
	Host struct {
		App *fiber.App
	}
)

func main() {
	// this is something that could be used in future project, when I want to use subdomain
	// trying out the subdomain
	hosts := map[string]*Host{}
	// adminWeb := AdminWeb{}
	// web := Web{}

	// testing admin
	admin := fiber.New()
	hosts["admin.localhost:8080"] = &Host{admin}

	// adminWeb.RegisterAdminHandlers(admin)

	admin.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"From": "from Admin"})
	})

	admin.Get("/:string", func(c *fiber.Ctx) error {
		result := c.Params("string")
		return c.SendString("From Admin - " + result)
	})

	// testing normal
	site := fiber.New()
	hosts["localhost:8080"] = &Host{site}

	// web.RegisterHandlers(site)

	site.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("From Normal")
	})

	site.Get("/:string", func(c *fiber.Ctx) error {
		result := c.Params("string")
		return c.SendString("From Normal - " + result)
	})

	// the routing
	e := fiber.New()
	e.Get("/*", func(c *fiber.Ctx) error {
		req := c.Request()
		// res := c.Response()
		host := hosts[string(req.Host())]

		if host == nil {
			return fiber.ErrNotFound
		} else {
			host.App.Server()
		}

		return nil
	})

	log.Fatal(e.Listen(":8080"))
}
