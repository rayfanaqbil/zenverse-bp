// url.go

package url

import (
    "github.com/rayfanaqbil/Zenverse-BP/controller"
    "github.com/rayfanaqbil/Zenverse-BP/handlers"
    "github.com/rayfanaqbil/Zenverse-BP/middleware"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/swagger"
    "go.mongodb.org/mongo-driver/mongo"
)

func Web(page *fiber.App, db *mongo.Database) {
    page.Get("/", controller.Sink)
    page.Post("/", controller.Sink)
    page.Put("/", controller.Sink)
    page.Patch("/", controller.Sink)
    page.Delete("/", controller.Sink)
    page.Options("/", controller.Sink)

    page.Get("/games", controller.GetAllGames)
    page.Get("/games/search", controller.GetGameByName)
    page.Get("/games/:id", controller.GetGamesByID)
    page.Put("/update/:id", controller.UpdateDataGames)
    page.Delete("/delete/:id", controller.DeleteGamesByID)
    page.Get("/docs/*", swagger.HandlerDefault)
    page.Post("/insert", controller.InsertDataGames)
    page.Post("/login", handlers.Login)
    page.Get("/admin", controller.GetDataAdmin)
	page.Post("/login/save-token", handlers.SaveToken)

    // Protected routes
    protected := page.Group("/login", middleware.Protected(db))
    protected.Get("/protected-route", func(c *fiber.Ctx) error {
        username := c.Locals("username")
        return c.JSON(fiber.Map{"message": "This is a protected route", "user": username})
    })
}
