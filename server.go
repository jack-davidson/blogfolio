package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var (
	Host       = "localhost"
	PasswdFile = "dbpasswd"
	Port       = 5432
	User       = "postgres"
	Password   = readPassword(PasswdFile)
	DB         = "test"
)

func readPassword(name string) string {
	content, _ := os.ReadFile(name)
	return string(content)
}

func main() {
	app := fiber.New()
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, DB)
	db, _ := sql.Open("postgres", conn)
	defer db.Close()

	app.Get("/", func(c *fiber.Ctx) error {
		/* Main Page */
		rows, err := db.Query("select name from colors where name = 'red';")
		defer rows.Close()
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			var color string
			rows.Scan(&color)
			fmt.Println(color)
		}
		return nil
	})

	app.Get("/blogs/:title", func(c *fiber.Ctx) error {
		/* Get Blog By Title */
		return nil
	})

	log.Fatal(app.Listen(":3000"))
}
