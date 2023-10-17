package router


import (
    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    app.Get("/", Home)
    app.Get("/about", About)
    app.Get("/contact", Contact)
}

func Home(c *fiber.Ctx) error {
    return c.SendFile("views/home.html")
}

func About(c *fiber.Ctx) error {
    return c.SendFile("views/about.html")
}

func Contact(c *fiber.Ctx) error {
    return c.SendFile("views/contact.html")
}