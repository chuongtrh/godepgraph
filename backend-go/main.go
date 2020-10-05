package main

import (
	"log"
	"os"

	"github.com/chuongtrh/godepviz/godep"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is ready!")
	})
	app.Get("/search/:pkg?", func(c *fiber.Ctx) error {

		pkgName := c.Query("pkg")
		node := &godep.Node{
			PkgName: pkgName,
			IsRoot:  true,
			Parent:  nil,
		}
		err := node.FindImports()
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		graph := node.BuildGraph()
		return c.SendString(graph)
	})
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	app.Listen(":" + port)
}
