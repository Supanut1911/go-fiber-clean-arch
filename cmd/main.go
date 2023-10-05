package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	app := fiber.New()

	//middleware
	app.Use(func(c *fiber.Ctx) error {
		//pass param to route
		c.Locals("name", "nutx")

		fmt.Printf("before")
		c.Next()
		fmt.Printf("after")
		return nil
	})

	app.Get("/", func(c *fiber.Ctx) error {
		name := c.Locals("name")
		fmt.Print("server up", name)
		return c.SendString("server up")
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("this post")
	})

	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString("name:" + name)
	})

	app.Get("/hello/:name/:sname", func(c *fiber.Ctx) error {
		name := c.Params("name")
		sname := c.Params("sname")
		return c.SendString("name:" + name + "sname" + sname)
	})

	app.Get("/num/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.ErrBadRequest
		}
		return c.SendString(fmt.Sprintf("ID = %v", id))
	})

	app.Get("/query", func(c *fiber.Ctx) error {
		name := c.Query("name")
		sname := c.Query("sname")
		return c.SendString("name:" + name + " " + sname)
	})

	app.Get("/query2", func(c *fiber.Ctx) error {
		person := Person{}
		c.QueryParser(&person)
		return c.JSON(person)
	})

	app.Listen(":3000")
}
