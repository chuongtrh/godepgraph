package main

import (
	"github.com/chuongtrh/godepviz/godep"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/:pkg?", func(c *fiber.Ctx) error {

		pkgName := c.Query("pkg")
		node := &godep.Node{
			PkgName: pkgName,
			IsRoot:  true,
			Parent:  nil,
		}
		node.FindImports()
		graph := node.BuildGraph()
		return c.SendString(graph)
	})
	app.Listen("localhost:3200")
}
