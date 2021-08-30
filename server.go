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

const (
	Intro = "Hello! I'm Jack Davidson, a student Software Engineer in Austin, Texas. I like playing around with gnu/linux machines and programming in Go."
)

type SiteContext struct {
	Title   string
	Name    string
	Email   string
	Content SiteContent
}

type SiteContent struct {
	Intro template.HTML /* Short introduction. */
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
		tmpl.Execute(c, SiteContext{
			Title: "jackdavidson.tech",
			Name:  "Jack Davidson",
			Email: "jackmaverickdavidson@gmail.com",
			Content: SiteContent{
				Intro: template.HTML(Intro),
			},
		})

		return nil
	})

	app.Get("/git/:repo", func(c *fiber.Ctx) error {
		c.Redirect(fmt.Sprintf("https://github.com/jack-davidson/%s", c.Params("repo")))
		return nil
	})

	log.Fatal(app.Listen(fmt.Sprintf("%s:%d", Host, Port)))
}
