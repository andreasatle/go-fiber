package routes

import (
	"github.com/andreasatle/go-fiber/src/handlers"
	"github.com/gofiber/fiber/v2"
)

func CreateRoutes(app *fiber.App) {
	public := app.Group("/public")
	public.Get("/home", handlers.GetPublicHome)
	public.Get("/resume", handlers.GetPublicResume)
	public.Get("/contact", handlers.GetPublicContact)

	auth := app.Group("/auth")
	auth.Get("/login", handlers.GetAuthLogin)
	auth.Post("/login", handlers.PostAuthLogin)
	auth.Get("/logout", handlers.GetAuthLogout)

	private := app.Group("/private")
	private.Get("/home", handlers.GetPrivateHome)
	app.Get("/", handlers.Redirect("/public/home"))
}