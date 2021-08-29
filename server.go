package main

import (
	"fmt"
	"log"
	"time"

	"html/template"

	"github.com/gofiber/fiber/v2"
)

const (
	Host = "0.0.0.0"
	Port = 3000
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

	app.Get("/git/:repo", func(c *fiber.Ctx) error {
		c.Redirect(fmt.Sprintf("https://github.com/jack-davidson/%s", c.Params("repo")))
		return nil
	})

	log.Fatal(app.Listen(fmt.Sprintf("%s:%d", Host, Port)))
}
