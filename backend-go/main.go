package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/chuongtrh/godepviz/godep"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/allegro/bigcache"
)

var (
	cache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(1440 * time.Minute))
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is ready!")
	})
	app.Get("/cache", func(c *fiber.Ctx) error {
		stats := cache.Stats()
		json, _ := json.Marshal(stats)
		return c.SendString(string(json))

	})
	app.Get("/search/:pkg?", func(c *fiber.Ctx) error {

		pkgName := c.Query("pkg")
		entry, errCached := cache.Get(pkgName)
		if errCached == nil {
			// fmt.Println("Cache found:", pkgName)
			return c.SendString(string(entry))
		}

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

		cache.Set(pkgName, []byte(graph))

		return c.SendString(graph)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "localhost:3200"
		// log.Fatal("$PORT must be set")
	} else {
		port = ":" + port
	}
	app.Listen(port)
}
