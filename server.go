package main

import (
	"fmt"
	"log"

	"html/template"

	"github.com/gofiber/fiber/v2"
)

type RootContext struct {
	Title string
}

func main() {
	app := fiber.New()
	app.Static("/static/", "./public")

	/* Main Page */
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(c)
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Fatal(err)
		}
		c.Context().SetContentType("text/html")
		tmpl.Execute(c, RootContext{
			Title: "jackdavidson",
		})

		return nil
	})

	/* Get Blog By Title */
	app.Get("/blogs/:title", func(c *fiber.Ctx) error {
		return nil
	})

	log.Fatal(app.Listen(":3000"))
}
