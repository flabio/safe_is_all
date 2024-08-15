package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/safe_msvc_city/insfratructure/routers"
)

func main() {
	app := fiber.New()
	// Custom CORS configuration

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080", // Specify allowed origins
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type, Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true, // Allow credentials
	}))

	routers.NewCityRouter(app)
	routers.NewStatesRouter(app)
	app.Listen(":3004")
}
