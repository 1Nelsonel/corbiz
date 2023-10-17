package main

import (
	"log"

	"github.com/1Nelsonel/corbiz/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("./views", ".html")

	
	app := fiber.New(fiber.Config{
			Views: engine,
	})

	// Serve static files (CSS, JavaScript, images) from the "public" directory
	app.Static("/static", "./public")

    // Set up application routes
    router.SetupRoutes(app)


	log.Fatal(app.Listen(":3000"))
}