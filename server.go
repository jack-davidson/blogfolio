package main

import (
	"fmt"
	"log"
	"time"

	"html/template"

	"github.com/gofiber/fiber/v2"
)

type IndexContext struct {
	Title string
	Name  string
}

func main() {
	app := fiber.New()
	app.Static("/", "./public", fiber.Static{
		CacheDuration: 1 * time.Microsecond,
	})

	/* Main Page */
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(c)
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Fatal(err)
		}
		c.Context().SetContentType("text/html")
		tmpl.Execute(c, IndexContext{
			Title: "jackdavidson.tech",
			Name:  "Jack Davidson",
		})

		return nil
	})

	/* Get Blog By Title */
	app.Get("/blogs/:title", func(c *fiber.Ctx) error {
		return nil
	})

	log.Fatal(app.Listen(":3000"))
}
