package main

import (
	"github.com/andreasatle/go-fiber/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/static", "./static")
	routes.CreateRoutes(app)
	app.ListenTLS(":8443", "cert.pem", "key.pem")
}